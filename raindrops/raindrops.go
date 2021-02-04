// Package raindrops takes an integer and returns a string of words that sound like rain based on its factors
package raindrops

import (
	"sort"
	"strconv"
	"strings"
)

var raindropWords = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert takes an integer and returns a string of raindrop sounds based on the factors defined in raindropWords
func Convert(input int) string {
	var raindropSound strings.Builder
	keys := sortedKeys()
	for _, key := range keys {
		if input%key == 0 {
			raindropSound.WriteString(raindropWords[key])
		}
	}
	// If there were no factors with raindrop sounds, return the number itself as a string
	if raindropSound.String() == "" {
		raindropSound.WriteString(strconv.Itoa(input))
	}
	return raindropSound.String()
}

func sortedKeys() (keys []int) {
	for key := range raindropWords {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return
}
