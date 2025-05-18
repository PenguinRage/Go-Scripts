package main

import (
	"fmt"
	"net/http"

	"stock-purchase-tracker/datatypes"
	"stock-purchase-tracker/utils"

	"github.com/gin-gonic/gin"
)

func validateValues(stock *datatypes.Stock) bool {
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
	var newStock datatypes.Stock

	if err := c.BindJSON(&newStock); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to BindJSON"})
		return
	}

	if !validateValues(&newStock) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed Validation Checks"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newStock)
}

func main() {
	router := gin.Default()
	router.POST("/stocks", buildStock)
	router.Run("localhost:8080")
}
