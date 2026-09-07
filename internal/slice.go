package internal

import (
	"fmt"
	"slices"
)

func InsertNumToSlice() []int {
	number := []int{50, 70, 66, 20, 32, 90}
	idx := slices.Index(number, 66) + 1
	leftSlices := make([]int, 3)
	copy(leftSlices, number[:idx])
	leftSlices = append(leftSlices, 88)
	newSlices := append(leftSlices, number[idx:]...)
	for _, v := range newSlices {
		fmt.Println(v)
	}
	return newSlices
}
