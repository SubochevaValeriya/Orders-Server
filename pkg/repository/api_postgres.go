package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"net/http"
	"time"
)

type ApiPostgres struct {
	db *sqlx.DB
}

func NewApiPostgres(db *sqlx.DB) *ApiPostgres {
	return &ApiPostgres{db: db}
}

func (r *ApiPostgres) CreateUser(user microservice.UsersBalances) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}
	var id int

	if user.Balance < 0 {
		return 0, errors.New("balance can't be negative")
	}
	changeBalanceQuery := fmt.Sprintf("INSERT INTO %s (balance) values ($1) RETURNING id", usersTable)
	row := r.db.QueryRow(changeBalanceQuery, user.Balance)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	addTransactionQuery := fmt.Sprintf("INSERT INTO %s (user_id, amount, reason, transaction_date) values ($1, $2, $3, $4)", transactionTable)
	_, err = tx.Exec(addTransactionQuery, id, user.Balance, ReasonOpening, time.Now())
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *ApiPostgres) GetBalanceById(userId int, ccy string) (microservice.UsersBalances, error) {
	// https://apilayer.com/marketplace/exchangerates_data-api?preview=true#documentation-tab
	var row microservice.UsersBalances
	var balance int
	getBalanceQuery := fmt.Sprintf("SELECT (balance) FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&balance, getBalanceQuery, userId)
	if err != nil {
		return row, err
	}

	if ccy != "" {

	} else {
		url := "https://api.apilayer.com/exchangerates_data/convert?to={to}&from={from}&amount={amount}"

		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set("apikey", "4EIJgz5GX9n5N8QUdRXHwQE01DfqGSqs")

		if err != nil {
			fmt.Println(err)
		}
		res, err := client.Do(req)
		if res.Body != nil {
			defer res.Body.Close()
		}
		body, err := ioutil.ReadAll(res.Body)

		fmt.Println(string(body))

		json.Unmarshal(responseData, &responseObject)
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)

	err := r.db.Get(&row, query, userId)

	return row, err
}
