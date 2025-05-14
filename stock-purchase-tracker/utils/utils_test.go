package utils

import (
	"testing"
)

func TestIsValidExchange(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid Exchange",
			input:    "NYSE",
			expected: true,
		},
		{
			name:     "Valid Exchange - Min Length",
			input:    "ASX",
			expected: true,
		},
		{
			name:     "Valid Exchange - Max Length",
			input:    "ABCDEF",
			expected: true,
		},
		{
			name:     "Invalid Exchange - Lowercase",
			input:    "nyse",
			expected: false,
		},
		{
			name:     "Invalid Exchange - Mixed Case",
			input:    "Nyse",
			expected: false,
		},
		{
			name:     "Invalid Exchange - Too Short",
			input:    "AB",
			expected: false,
		},
		{
			name:     "Invalid Exchange - Too Long",
			input:    "ABCDEFG",
			expected: false,
		},
		{
			name:     "Invalid Exchange - Empty String",
			input:    "",
			expected: false,
		},
		{
			name:     "Invalid Exchange - Special Characters",
			input:    "NYSE!",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsValidExchange(tc.input)
			if actual != tc.expected {
				t.Errorf("IsValidExchange(%s) = %v, expected %v", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestIsValidState(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid State - SELL",
			input:    "SELL",
			expected: true,
		},
		{
			name:     "Valid State - BUY",
			input:    "BUY",
			expected: true,
		},
		{
			name:     "Invalid State - Lowercase",
			input:    "sell",
			expected: false,
		},
		{
			name:     "Invalid State - Mixed Case",
			input:    "Sell",
			expected: false,
		},
		{
			name:     "Invalid State - Other Value",
			input:    "HOLD",
			expected: false,
		},
		{
			name:     "Invalid State - Empty String",
			input:    "",
			expected: false,
		},
		{
			name:     "Invalid State - Special Characters",
			input:    "SELL!",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsValidState(tc.input)
			if actual != tc.expected {
				t.Errorf("IsValidState(%s) = %v, expected %v", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestIsValidStockCode(t *testing.T) {
	tests := []struct {
		rule      string
		testValue string
		expected  bool
	}{
		{"Case 1: Valid Stock", "MSFT", true},
		{"Case 2: Invalid Stock - Lowercase", "msft", false},
		{"Case 3: Invalid Stock - Symbols", "BRK.B", true},
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
func TestIsAllAlphaCaps(t *testing.T) {
	tests := []struct {
		rule      string
		testValue string
		expected  bool
	}{
		{"Case 1: AllCaps", "ABSC", true},
		{"Case 2: Fails AllCaps", "AbCD", false},
		{"Case 3: Fails AllCaps - Symbols", "ABCD.", false},
		{"Case 4: Empty String", "", true},               // Empty string should return true
		{"Case 5: Numbers", "1234", false},               // Alpha only
		{"Case 6: Mixed Alphanumeric", "A1B2", false},    // No mixed // Alpha only
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
