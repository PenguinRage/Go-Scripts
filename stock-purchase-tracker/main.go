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
	if !utils.IsValidExchange(arr[0]) {
		fmt.Println("Invalid Exchange")
		return false
	} else if !utils.IsValidStockCodeBasic(arr[1]) {
		fmt.Println("Invalid Stock")
		return false
	} else if !utils.IsValidState(arr[2]) {
		fmt.Println("Invalid State")
		return false
	} else if !utils.IsValidQuantity(arr[3]) {
		fmt.Println("Invalid Quantity")
		return false
	} else if !utils.IsValidCurrency(arr[4]) {
		fmt.Println("Invalid Currency")
		return false
	} else if !utils.IsvalidPrice(arr[5]) {
		fmt.Println("Invalid Price")
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
