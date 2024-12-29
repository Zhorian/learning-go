package main

import (
	"fmt"
	"strconv"
	"strings"
)

type score struct {
	name  string
	score int
}

func main() {
	scores := []score{
		{"Dent, Arthur", 87},
		{"MacMillian, Tricia", 96},
		{"Prefect, Ford", 64},
	}

	fmt.Println("Select score to print (1 - 3):")
	var option string
	fmt.Scanln(&option)
	optionAsInt, err := strconv.Atoi(option)

	if err != nil {
		fmt.Println("Could not covert option to a number!")
		return
	}

	index := 0
	// if option == "1" {
	// 	index = 0
	// } else if option == "2" {
	// 	index = 1
	// } else if option == "3" {
	// 	index = 2
	// } else {
	// 	fmt.Printf("Unknown option, using default of %d", index+1)
	// }

	switch optionAsInt {
	case 1, 2, 3:
		index = optionAsInt - 1
	default:
		fmt.Printf("Unknown option, using default of %d", index+1)
	}

	fmt.Println("Student scores")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(scores[index].name, scores[index].score)
}
