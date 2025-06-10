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
		{"Case 1: Valid intake", &datatypes.Stock{"ASX", "ASX", "BUY", 1, "AUD", 1.00}, true},
		{"Case 2: Valid intake", &datatypes.Stock{"NYSE", "BRK.B", "BUY", 1000, "USD", 512.15}, true},
		{"Case 3: Valid intake", &datatypes.Stock{"NYSE", "V", "SELL", 1000, "USD", 366.66}, true},
	}

	for _, tt := range tests {
		result := validateValues(tt.testValue)
		if tt.expected != result {
			t.Errorf("%v - expected=%v result=%v", tt.rule, tt.expected, result)
		}
	}
}
