// Package proverb takes a slice of words and modifies the proverb "For want of a nail.." to use the provided words
package proverb

import "fmt"

const (
	line = "For want of a %s the %s was lost."
	last = "And all for the want of a %s."
)

// Proverb takes a slice of words and interjects them into the proverb in the format:
//   "For want of a <...> the <...> was lost."
//   ...
//   "And all for the want of a <...>."
func Proverb(rhyme []string) []string {
	numberOfLines := len(rhyme)
	if numberOfLines == 0 {
		return []string{}
	}
	// Add last line
	var proverb []string = make([]string, numberOfLines)
	proverb[numberOfLines-1] = fmt.Sprintf(last, rhyme[0])

	if numberOfLines == 1 {
		return proverb
	}

	// Add beginning lines
	for i := 0; i < numberOfLines-1; i++ {
		proverb[i] = fmt.Sprintf(line, rhyme[i], rhyme[i+1])
	}

	return proverb
}
