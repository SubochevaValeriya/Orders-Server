package repository

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	order "http_server"
	"time"
)

type ApiPostgres struct {
	db    *sqlx.DB
	cache *redis.Client
}

func NewApiPostgres(db *sqlx.DB, cache *redis.Client) *ApiPostgres {
	return &ApiPostgres{
		db:    db,
		cache: cache}
}

//CreateOrder adds order to DB and to cache
func (r *ApiPostgres) CreateOrder(order order.Order) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}
	var id int

	changeBalanceQuery := fmt.Sprintf("INSERT INTO %s (data) values ($1) RETURNING order_id", ordersTable)
	jsonOrder, err := json.Marshal(order)
	if err != nil {
		return 0, err
	}
	row := r.db.QueryRow(changeBalanceQuery, jsonOrder)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		value, err := json.Marshal(order)
		if err != nil {
			return id, err
		}
		set := r.cache.Set(string(id), value, 1000*time.Minute)
		if _, err := set.Result(); err == nil {
			logrus.Println("order was added to the cache")
		}
	}

	return id, err
}

//GetOrderById gets order from cache or from db (and add to cache)
func (r *ApiPostgres) GetOrderById(orderId int) (order.Order, error) {
	var order order.Order
	orderCash, err := r.cache.Get(string(orderId)).Bytes()
	if err == nil {
		logrus.Println("order was found in the cache")
		err := json.Unmarshal(orderCash, &order)
		return order, err
	}

	var row string
	query := fmt.Sprintf("SELECT (data) FROM %s WHERE order_id=$1", ordersTable)
	err = r.db.Get(&row, query, orderId)
	if err != nil {
		return order, fmt.Errorf("error while trying to get order by Id from DB: %w", err)
	}

	err = json.Unmarshal([]byte(row), &order)
	set := r.cache.Set(string(orderId), row, 1000*time.Minute)
	if _, err := set.Result(); err == nil {
		logrus.Println("order was added to the cache")
	}

	return order, err
}
