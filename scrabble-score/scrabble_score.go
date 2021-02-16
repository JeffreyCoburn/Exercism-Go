// Package scrabble computes the score for a word
package scrabble

import (
	"errors"
)

/* From Scrabble rules:
Letter                           Value
A, E, I, O, U, L, N, R, S, T       1
D, G                               2
B, C, M, P                         3
F, H, V, W, Y                      4
K                                  5
J, X                               8
Q, Z                               10
*/
var letterScores = map[rune]int{
	'A': 1, 'a': 1,
	'E': 1, 'e': 1,
	'I': 1, 'i': 1,
	'O': 1, 'o': 1,
	'U': 1, 'u': 1,
	'L': 1, 'l': 1,
	'N': 1, 'n': 1,
	'R': 1, 'r': 1,
	'S': 1, 's': 1,
	'T': 1, 't': 1,
	'D': 2, 'd': 2,
	'G': 2, 'g': 2,
	'B': 3, 'b': 3,
	'C': 3, 'c': 3,
	'M': 3, 'm': 3,
	'P': 3, 'p': 3,
	'F': 4, 'f': 4,
	'H': 4, 'h': 4,
	'V': 4, 'v': 4,
	'W': 4, 'w': 4,
	'Y': 4, 'y': 4,
	'K': 5, 'k': 5,
	'J': 8, 'j': 8,
	'X': 8, 'x': 8,
	'Q': 10, 'q': 10,
	'Z': 10, 'z': 10,
}

// Word and letter modifiers
const (
	single = 1
	double = 2
	triple = 3
)

// ErrInvalidLength is returned if the number of letter modifiers does not match the number of letters
var ErrInvalidLength = errors.New("Must have exactly one letter modifier for each letter in the word")

// ErrInvalidModifier is returned if the letter or word modifier is not valid
var ErrInvalidModifier = errors.New("Invalid modifier.  Must be single(x1), double(x2), or triple(x3)")

// Score calculates the score of a scrabble word with no special modifiers, such as double/triple letters or words
func Score(word string) (score int) {
	for _, letter := range word {
		score += letterScores[letter]
	}
	return
}

// ExtendedScore calculates the score of a scrabble word given the word itself and an array containing the letter modifiers as well as the word modifier.
func ExtendedScore(word string, letterModifiers []int, wordModifier int) (int, error) {
	if wordModifier < single || wordModifier > triple {
		return 0, ErrInvalidModifier
	}
	var score int
	if len(word) != len(letterModifiers) {
		return 0, ErrInvalidLength
	}
	for i, letter := range word {
		if letterModifiers[i] < single || letterModifiers[i] > triple {
			return 0, ErrInvalidModifier
		}
		score += letterModifiers[i] * letterScores[letter]
	}
	score *= wordModifier
	return score, nil
}
