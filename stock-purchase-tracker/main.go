package main

import (
	"fmt"
	"os"

	"stock-purchase-tracker/utils"
)

type Stock struct {
	Exchange string  `db:"exchange"`
	Code     string  `db:"code"`
	State    string  `db:"state"`
	Quantity int     `db:"quantity"`
	Currency string  `db:"currency"`
	Price    float64 `db:"price"`
}

func validateValues(arr []string) bool {
	check := true

	check = utils.IsValidExchange(arr[0])

	if check == false {
		return false
	}

	return true
}

// buildStock function  
func buildStock(exchange string, code string, state string, quantity int, curency string, price float64) *Stock {
	stock := Stock{exchange, code, state, quantity, curency, price}
	return &stock
}

func main() {
	args := os.Args[1:]
	validated := validateValues(args)

	fmt.Println(validated)
}
