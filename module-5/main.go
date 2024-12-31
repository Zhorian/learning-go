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

func getOption() (option string) {
	fmt.Println("1) Enter Score")
	fmt.Println("2) Print Report")
	fmt.Println("q) Quit")
	fmt.Println()
	fmt.Println("Select an option")

	fmt.Scanln(&option)

	return option
}

func addScore(scores *[]score) {
	fmt.Println("Enter a student name and score")
	var name string
	var rawScore string
	fmt.Scanln(&name, &rawScore)

	var s, err = strconv.Atoi(rawScore)
	if err != nil {
		fmt.Println("Could not covert option to a number!")
		return
	}

	*scores = append(*scores, score{name: name, score: s})
}

func printReport(scores *[]score) {
	fmt.Println("Student scores")
	fmt.Println(strings.Repeat("-", 14))

	for i, s := range *scores {
		fmt.Println(i, s.name, s.score)
	}

	fmt.Println(strings.Repeat("-", 14))
}

func main() {
	scores := []score{}
	shouldContinue := true

	for shouldContinue {
		option := getOption()

		switch option {
		case "1":
			addScore(&scores)
		case "2":
			printReport(&scores)
		case "q":
			shouldContinue = false
		}
	}
}
