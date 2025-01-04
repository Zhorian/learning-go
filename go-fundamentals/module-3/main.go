package main

import (
	"fmt"
	"strings"
)

type score struct {
	name  string
	score int
}

func main() {
	// name, score := "Dent, Arthur", 87
	// students := []string{"Dent, Arthur",
	// 	"MacMillan, Tricia",
	// 	"Prefect, Ford",
	// }

	// scores := []int{87, 96, 64}
	// scores := map[string]int{
	// 	students[0]: 87,
	// 	students[1]: 96,
	// 	students[2]: 64,
	// }

	scores := []score{
		{"Dent, Arthur", 87},
		{"MacMillian, Tricia", 96},
		{"Prefect, Ford", 64},
	}

	fmt.Println("Student scores")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(scores[0].name, scores[0].score)
	fmt.Println(scores[1].name, scores[1].score)
	fmt.Println(scores[2].name, scores[2].score)
}
