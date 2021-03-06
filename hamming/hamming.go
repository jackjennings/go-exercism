// Package hamming implements functions for working with hamming scores.
package hamming

import "errors"

// Distance returns the number of nucleotides that differ between two
// DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("DNA strands are of unequal length")
	}

	distance := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
