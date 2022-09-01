package repository

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	order "http_server"
	"time"
)

type ApiPostgres struct {
	db   *sqlx.DB
	cash *redis.Client
}

func NewApiPostgres(db *sqlx.DB, cash *redis.Client) *ApiPostgres {
	return &ApiPostgres{
		db:   db,
		cash: cash}
}

func (r *ApiPostgres) CreateOrder(order order.Order) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}
	var id int

	fmt.Println("from cr")
	changeBalanceQuery := fmt.Sprintf("INSERT INTO %s (data) values ($1) RETURNING order_id", ordersTable)
	jsonOrder, err := json.Marshal(order)
	if err != nil {
		return 0, err
	}
	row := r.db.QueryRow(changeBalanceQuery, jsonOrder)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		fmt.Println(err)
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		r.cash.Set(string(id), order, 1000*time.Minute)
	}

	fmt.Println("from cr2")
	return id, tx.Commit()
}

func (r *ApiPostgres) GetOrderById(orderId int) (order.Order, error) {
	var order order.Order
	orderCash, err := r.cash.Get("orderId").Bytes()
	if err == nil {
		fmt.Println("we found!")
		err := json.Unmarshal(orderCash, &order)
		return order, err
	}

	fmt.Println("not found((")
	var row string
	query := fmt.Sprintf("SELECT (data) FROM %s WHERE order_id=$1", ordersTable)
	err = r.db.Get(&row, query, orderId)
	if err != nil {
		return order, fmt.Errorf("error while trying to get order by Id from DB: %w", err)
	}

	err = json.Unmarshal([]byte(row), &order)
	p := r.cash.Set(string(orderId), row, 1000*time.Minute)
	fmt.Println(p.Result())
	//.Err() != nil {
	//	return order, fmt.Errorf("can't add order to Cash: %w", err)
	//}

	return order, err

	//if data, ok := m[orderId]; ok{
	//	return data, nil
	//}
	//
	//var row order.Order
	//query := fmt.Sprintf("SELECT (data) FROM %s WHERE order_id=$1", ordersTable)
	//err = r.db.Get(&row, query, orderId)
	//if err != nil {
	//	return order, fmt.Errorf("error while trying to get order by Id from DB: %w", err)
	//}
	//err = json.Unmarshal([]byte(row), &order)
	//return order, err
}
