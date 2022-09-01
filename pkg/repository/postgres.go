package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	ordersTable = "orders"
)

type ConfigPostgres struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg ConfigPostgres) (*sqlx.DB, error) {
	//url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
	//	cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	//db, err := sqlx.Connect("postgres", url)
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, fmt.Errorf("connect error: %s", err)
	}

	return db, nil
}
