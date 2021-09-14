package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram checks if a string is na isogram
func IsIsogram(input string) bool {
	charsInInput := make(map[rune]int) // Stores the characters of the input string
	input = strings.ToLower(input)
	for _, char := range input {
		if !unicode.IsLetter(char) {
			continue
		}
		// Check if character has been used already
		_, found := charsInInput[char]
		if found {
			return false
		}
		// If the character wasn't found, add the character to the map
		charsInInput[char] = 1
	}
	return true
}
