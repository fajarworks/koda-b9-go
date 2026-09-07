package internal

import "fmt"

func Window(x int, y int) error {
	var a string
	var b string
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if i == 0 || i == x-1 || j == 0 || j == y-1 {
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
