package main

import (
	"testing"
)

func TestValidateValues(t *testing.T) {
	tests := []struct {
		rule      string
		testValue []string
		expected  bool
	}{
		{"Case 1: Valid intake", []string{"ASX", "ASX", "BUY", "1", "AUD", "1.00"}, true},
		{"Case 2: Valid intake", []string{"NADAQ", "BRK.B", "BUY", "1000", "USD", "512.15"}, true},
	}

	for _, tt := range tests {
		result := validateValues(tt.testValue)
		if tt.expected != result {
			t.Errorf("%v - expected=%v result=%v", tt.rule, tt.expected, result)
		}
	}
}
