package repository

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	order "http_server"
)

type ApiPostgres struct {
	db *sqlx.DB
}

func NewApiPostgres(db *sqlx.DB) *ApiPostgres {
	return &ApiPostgres{db: db}
}

func newCashMap() map[int]order.Order {
	return map[int]order.Order{}
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

	//m[id] = order
	fmt.Println("from cr2")
	return id, tx.Commit()
}

func (r *ApiPostgres) GetOrderById(orderId int) (order.Order, error) {
	//if data, ok := m[orderId]; ok{
	//	return data, nil
	//}

	//var row order.Order
	var row string
	query := fmt.Sprintf("SELECT (data) FROM %s WHERE order_id=$1", ordersTable)
	err := r.db.Get(&row, query, orderId)
	var order order.Order
	json.Unmarshal([]byte(row), &order)
	if err != nil {
		return order, fmt.Errorf("error while trying to get order by Id from DB: %w", err)
	}

	return order, err
}
