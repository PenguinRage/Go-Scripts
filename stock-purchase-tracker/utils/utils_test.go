package utils

import (
	"testing"
)

func TestIsValidStockCode(t *testing.T) {
	tests := []struct {
		rule      string
		testValue string
		expected  bool
	}{
		{"Case 1: Valid Stock", "MSFT", true},
		{"Case 2: Invalid Stock - Lowercase", "msft", false},
		{"Case 3: Invalid Stock - Symbols", "MSFT.", false},
		{"Case 4: Invalid Stock - Empty", "", false},
		{"Case 5: Invalid Stock - Number", "1234", false},
	}

	for _, tt := range tests {
		result := IsValidStockCodeBasic(tt.testValue)
		if tt.expected != result {
			t.Errorf("%v - expected=%v result=%v", tt.rule, tt.expected, result)
		}
	}
}

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
		{"Case 4: Empty String", "", true},               // Empty string should return true
		{"Case 5: Numbers", "1234", true},                // Numbers should return true
		{"Case 6: Mixed Alphanumeric", "A1B2", true},     // Mixed alphanumeric should return true
		{"Case 7: Lowercase and Numbers", "a1b2", false}, // Lowercase and numbers should return false
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
		{"Case 4: Empty String Valid", "", 0, 5, true},          // Empty string within valid range
		{"Case 5: Empty String Invalid", "", 1, 5, false},       // Empty string outside valid range
		{"Case 6: Min equals Max Valid", "ABC", 3, 3, true},     // Min equals max and is valid
		{"Case 7: Min equals Max Invalid", "ABCD", 3, 3, false}, // Min equals max and is invalid
	}

	for _, tt := range tests {
		result := IsValidLength(tt.testValue, tt.min, tt.max)
		if tt.expected != result {
			t.Errorf("%v - expected=%v result=%v", tt.rule, tt.expected, result)
		}
	}

}
