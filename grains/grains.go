// Package grains implements functions relating to chess-based wheat cultivation
package grains

import "fmt"

// Square returns the number of wheat grains on the nth square of a chess board
func Square(index int) (uint64, error) {
	if index <= 0 || index > 64 {
		return 0, fmt.Errorf("Square number %d is invalid", index)
	}

	return 1 << uint64(index-1), nil
}

// Total returns the number of wheat grains on a (very large) chess board
func Total() uint64 {
	return uint64(1<<64 - 1)
}
