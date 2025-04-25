package main

import (
	"os"
)

type Stock struct {
	Code     string  `db:"code"`
	State    string  `db:"state"`
	Quantity int64   `db:"quantity"`
	Price    float64 `db:"price"`
}

func validateValues(arr []string) {

}

// buildStock function  
func buildStock(code string, state string, quantity int64, price float64) *Stock {
	stock := Stock{code, state, quantity, price}
	return &stock
}

func main() {
	args := os.Args[1:]
	validateValues(args)

}
