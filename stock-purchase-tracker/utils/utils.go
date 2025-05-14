package utils

import (
	"regexp"
	"unicode"
	"unicode/utf8"
)

func IsAllCaps(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func IsValidLength(s string, minLength int, maxLength int) bool {
	length := utf8.RuneCountInString(s)

	if minLength != -1 && length < minLength {
		return false
	}

	if maxLength != -1 && length > maxLength {
		return false
	}

	return true
}

func IsValidExchange(s string) bool {
	return IsAllCaps(s) && IsValidLength(s, 3, 6)
}

func IsValidState(s string) bool {
	if s == "SELL" || s == "BUY" {
		return true
	}
	return false
}

func IsValidCurrency(s string) bool {
	if s == "AUD" || s == "USD" {
		return true
	}
	return false
}

func IsValidStockCodeBasic(code string) bool {
	if len(code) < 1 || len(code) > 6 { // Increased max length to accommodate BRK.B
		return false
	}
	return regexp.MustCompile(`^[A-Z]+\.?[A-Z]*$`).MatchString(code)
}
