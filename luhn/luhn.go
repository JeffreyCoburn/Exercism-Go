package luhn

import (
	"strings"
	"unicode"
)

func Valid(input string) bool {
	// Check for non-valid characters and remove white space (removing all whitespace even though instructions only mention spaces)
	var b strings.Builder
	b.Grow(len(input))
	for _, c := range input {
		if !unicode.IsSpace(c) && !unicode.IsDigit(c) {
			return false
		}
		if unicode.IsDigit(c) {
			b.WriteRune(c)
		}
	}
	digits := b.String()

	// Check for single digit numbers
	if len(digits) < 2 {
		return false
	}

	// Sum the digits, doubling every 2nd digit from the right and subtracting 9 if the result is > 9
	var sumOfDigits int
	for i, c := range digits {
		digit := int(c - '0')
		if i%2 == len(digits)%2 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sumOfDigits += digit
	}

	return sumOfDigits%10 == 0
}
