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
	if len(rhyme) == 0 {
		return []string{}
	}
	var proverb []string = make([]string, len(rhyme))
	proverb[len(rhyme)-1] = fmt.Sprintf(last, rhyme[0])
	if len(rhyme) == 1 {
		return proverb
	}
	for i := 0; i < len(rhyme)-1; i++ {
		proverb[i] = fmt.Sprintf(line, rhyme[i], rhyme[i+1])
	}
	return proverb
}
