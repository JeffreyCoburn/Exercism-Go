// Package hamming calculates the Hamming distance between two DNA strand sequences
package hamming

import "errors"

var ErrDifferentLengths = errors.New("cannot compare DNA strands of different lengths")

// Distance returns the hamming distance between two DNA sequences
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, ErrDifferentLengths
	}
	hammingDistance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
