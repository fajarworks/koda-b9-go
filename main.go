package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fajarworks/koda-b9-go/internal"
	"github.com/fajarworks/koda-b9-go/internal/model"
	"github.com/fajarworks/koda-b9-go/internal/rectangle"
	"github.com/fajarworks/koda-b9-go/pkg"
)

func main() {
	fmt.Println("======== MINI TASK GOLANG DAY 1 ========")
	fmt.Println()
	fmt.Println("1. Area of Rectangle")
	fmt.Println("2. Circumference of Rectangle")
	fmt.Println("3. Area and Circumference of Rectangle")
	fmt.Println("4. Make a window with loop")
	fmt.Println("5. Insert a 88 after 66 in slices")
	fmt.Println("6. Create biodata from struct field")
	fmt.Println("7. Read and File")
	fmt.Println("8. Get Person From Struct")
	fmt.Println("0. Exit")
	fmt.Println()
	fmt.Println("========================================")

	scanner := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("\nChoose menu: ")

		text, _ := scanner.ReadString('\n')
		text = strings.TrimSpace(text)

		switch text {
		case "1":
			fmt.Println(rectangle.AreaOfRectangle(10, 10))

		case "2":
			fmt.Println(rectangle.CircumferenceOfRectangle(10, 10))

		case "3":
			area, circumference := rectangle.AreaAndCircumOfRectangle(10, 10)
			fmt.Println("Area:", area)
			fmt.Println("Circumference:", circumference)

		case "4":
			internal.Window(5, 10)

		case "5":
			internal.InsertNumToSlice()

		case "6":
			model.GetBio()

		case "7":
			res, err := pkg.OpenAndReadFile(".\\readme.md")
			fmt.Println(err)
			fmt.Print(res)

		case "8":

			person := model.NewPerson("maaruf", "palopo", "081234567890")
			result := person.GetPersonData()
			fmt.Println(result)
			greet := person.Greet()
			fmt.Println(greet)

			person.SetPersonName("hidayat")

			result = person.GetPersonData()
			fmt.Println(result)
			greet = person.Greet()
			fmt.Println(greet)

		case "0":
			fmt.Println("Goodbye!")
			os.Exit(0)

		default:
			fmt.Println("Invalid menu")
		}
	}

}
