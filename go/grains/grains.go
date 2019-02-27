package grains

import "errors"

// Square returns the number of grains on a given square
func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 0, errors.New("square must be in range [1,64]")
	}
	grains := 1
	for i := 1; i < square; i++ {
		grains *= 2
	}
	return uint64(grains), nil
}

// Total returns the total number of grains on the board
func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		n, _ := Square(i)
		total += n
	}
	return total
}
