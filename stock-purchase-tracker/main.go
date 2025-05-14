package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

func validateValues(stock *Stock) bool {
	if !utils.IsValidExchange(stock.Exchange) {
		fmt.Println("Invalid Exchange")
		return false
	} else if !utils.IsValidStockCodeBasic(stock.Code) {
		fmt.Println("Invalid Stock")
		return false
	} else if !utils.IsValidState(stock.State) {
		fmt.Println("Invalid State")
		return false
	} else if !utils.IsValidCurrency(stock.Currency) {
		fmt.Println("Invalid Currency")
		return false
	}
	return true
}

// buildStock function  
func buildStock(c *gin.Context) {
	var newStock Stock

	if err := c.BindJSON(&newStock); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Values"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newStock)
}

func main() {
	router := gin.Default()
	router.POST("/stocks", buildStock)
	router.Run("localhost:8080")
}
