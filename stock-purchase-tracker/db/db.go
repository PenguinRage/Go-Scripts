package db

import (
	"log"
	"stock-purchase-tracker/datatypes"

	"github.com/jmoiron/sqlx"
)

func InsertStock(stock *datatypes.Stock) {
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.Close()

}
