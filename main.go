package main

import (
	"fmt"
	"slices"

	"github.com/fajarworks/koda-b9-go/internal/model"
)

func areaOfRectangle(p int8, l int8) int16 {
	return int16(p) * int16(l)

}

func circumferenceOfRectangle(p int8, l int8) int16 {
	return 2 * (int16(p) + int16(l))
}

func areaAndCircumOfRectangle(p int8, l int8) (area int16, circum int16) {
	area = int16(p) * int16(l)
	circum = 2 * (int16(p) + int16(l))
	return area, circum
}

func getAreaAndCircum(area func(int8, int8) int16, circum func(int8, int8) int16, p int8, l int8) (a int16, b int16) {
	a = area(p, l)
	b = circum(p, l)
	return a, b
}

func window(x int) error {
	var a string
	var b string
	for i := 0; i < x; i++ {
		for j := 0; j < x; j++ {
			if i == 0 || i == x-1 || j == 0 || j == x-1 {
				a = "* "
				fmt.Print(a)
			} else {
				a = "  "
				fmt.Print(a)
			}
		}
		b = " "
		fmt.Println(b)

	}
	return nil
}

func insertNumToSlice() []int {
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

func main() {
	var p int8 = 5
	var l int8 = 10

	fmt.Println(areaOfRectangle(p, l))
	fmt.Println(circumferenceOfRectangle(p, l))

	area, circum := areaAndCircumOfRectangle(p, l)
	fmt.Printf("area of rectangle: %d\n", area)
	fmt.Printf("circumference of rectangle: %d \n", circum)

	window(3)
	insertNumToSlice()
	model.GetBio()
}
