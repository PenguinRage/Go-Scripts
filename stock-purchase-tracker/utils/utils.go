package utils

import (
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
	return IsAllCaps(s) && IsValidLength(s, 2, 4)
}

func IsValidState(s string) bool {
	if s == "SELL" || s == "BUY" {
		return true
	}
	return false
}

func IsValidStock(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) || r != '.' {
			return false
		}
	}
	return true
}

