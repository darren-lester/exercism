package grains

import "errors"

// Square returns the number of grains on a given square
func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 0, errors.New("square must be in range [1,64]")
	}
	return 1 << uint(square-1), nil
}

// Total returns the total number of grains on the board
func Total() uint64 {
	return 1<<64 - 1
}
