package grains

import (
	"errors"
)

const numberOfSquares = 64

const totalGrains uint64 = 1<<numberOfSquares - 1

func Square(square int) (uint64, error) {
	if square < 1 || square > numberOfSquares {
		return 0, errors.New("not a valid square")
	}
	var numOfGrains uint64 = 1 << (square - 1)
	return numOfGrains, nil
}

func Total() uint64 {
	return totalGrains
}
