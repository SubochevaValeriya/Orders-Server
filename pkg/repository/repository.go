package repository

import (
	"github.com/jmoiron/sqlx"
)

type Balance interface {
	CreateUser(user microservice.UsersBalances) (int, error)
	GetBalanceById(userId int) (microservice.UsersBalances, error)
}

type Repository struct {
	Balance
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{NewApiPostgres(db)}
}
