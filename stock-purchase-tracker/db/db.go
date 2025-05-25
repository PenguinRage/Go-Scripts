package db

import (
	"log"
	"stock-purchase-tracker/datatypes"

	"github.com/jmoiron/sqlx"
)

func InsertStock(stock *datatypes.Stock) {
	db, err := sqlx.Connect("postgres", "user=icleasby dbname=stocks sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	db.MustExec("INSERT INTO stocks (exchange, code, state, quantity, purchase_price) VALUES ($1, $2, $3, $4, $5)", stock.Exchange, stock.Code, stock.State, stock.Quantity, stock.Price)

	db.Close()

}
