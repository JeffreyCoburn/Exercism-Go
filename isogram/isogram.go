package isogram

import "unicode"

// IsIsogram checks if a string is na isogram
func IsIsogram(input string) bool {
	charsInInput := make(map[rune]int) // Stores the characters of the input string
	for _, char := range input {
		// Check if character has been used already
		_, found := charsInInput[unicode.ToLower(char)]
		if found {
			return false
		}
		// If the character wasn't found, add the character to the map
		if unicode.IsLetter(char) {
			charsInInput[unicode.ToLower(char)] = 1
		}
	}
	return true

}
