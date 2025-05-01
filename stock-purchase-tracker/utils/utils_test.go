package utils

import (
	"testing"
)

// TestIsAllCaps function  
func TestIsAllCaps(t *testing.T) {
	tests := []struct {
		rule      string
		testValue string
		expected  bool
	}{
		{"Case 1: AllCaps", "ABSC", true},
		{"Case 2: Fails AllCaps", "AbCD", false},
		{"Case 3: Fails AllCaps - Symbols", "ABCD.", false},
	}

	for _, tt := range tests {
		result := IsAllCaps(tt.testValue)
		if tt.expected != result {
			t.Errorf("%v - expected=%v result=%v", tt.rule, tt.expected, result)
		}
	}
}

func TestIsValidLength(t *testing.T) {
	tests := []struct {
		rule      string
		testValue string
		min       int
		max       int
		expected  bool
	}{
		{"Case 1: Valid Size", "ABCDE", 2, 5, true},
		{"Case 2: Invalid Min", "A", 2, 5, false},
		{"Case 3: Invalid Max", "ABCDEFG", 2, 6, false},
	}

	for _, tt := range tests {
		result := IsValidLength(tt.testValue, tt.min, tt.max)
		if tt.expected != result {
			t.Errorf("%v - expected=%v result=%v", tt.rule, tt.expected, result)
		}
	}

}
