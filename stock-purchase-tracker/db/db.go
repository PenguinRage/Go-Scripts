package db

import (
	"log"
	"stock-purchase-tracker/datatypes"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InsertStock(stock *datatypes.Stock) {
	db, err := sqlx.Connect("postgres", "user=stockuser password=stockpass dbname=stocks host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	db.MustExec("INSERT INTO stocks (exchange, code, state, quantity, currency, price) VALUES ($1, $2, $3, $4, $5, $6)", stock.Exchange, stock.Code, stock.State, stock.Quantity, stock.Currency, stock.Price)

	db.Close()

}

func GetStocks() []datatypes.Stock {
	db, err := sqlx.Connect("postgres", "user=stockuser password=stockpass dbname=stocks host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	var stocks []datatypes.Stock
	err = db.Select(&stocks, "SELECT exchange, code, state, quantity, currency, price FROM stocks")
	if err != nil {
		log.Printf("Error fetching stocks: %v", err)
		return []datatypes.Stock{}
	}

	return stocks
}
