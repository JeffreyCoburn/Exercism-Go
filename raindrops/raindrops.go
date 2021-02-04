// Package raindrops takes an integer and returns a string of words that sound like rain based on its factors
package raindrops

import (
	"sort"
	"strconv"
)

var raindropWords = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert takes an integer and returns a string of raindrop sounds based on its factors
func Convert(input int) (raindropSound string) {
	keys := sortedKeys()
	for _, key := range keys {
		if input%key == 0 {
			raindropSound += raindropWords[key]
		}
	}
	if raindropSound == "" {
		raindropSound = strconv.Itoa(input)
	}
	return raindropSound
}

func sortedKeys() (keys []int) {
	for key := range raindropWords {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return
}
