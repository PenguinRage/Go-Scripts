package main

import (
	"stock-purchase-tracker/datatypes"
	"testing"
)

func TestValidateValues(t *testing.T) {
	tests := []struct {
		rule      string
		testValue *datatypes.Stock
		expected  bool
	}{
		{"Case 1: Valid intake", &datatypes.Stock{Exchange: "ASX", Code: "ASX", State: "BUY", Quantity: 1, Currency: "AUD", Price: 1.00}, true},
		{"Case 2: Valid intake", &datatypes.Stock{Exchange: "NYSE", Code: "BRK.B", State: "BUY", Quantity: 1000, Currency: "USD", Price: 512.15}, true},
		{"Case 3: Valid intake", &datatypes.Stock{Exchange: "NYSE", Code: "V", State: "SELL", Quantity: 1000, Currency: "USD", Price: 366.66}, true},
	}

	for _, tt := range tests {
		result := validateValues(tt.testValue)
		if tt.expected != result {
			t.Errorf("%v - expected=%v result=%v", tt.rule, tt.expected, result)
		}
	}
}
