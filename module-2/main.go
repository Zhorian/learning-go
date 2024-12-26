package main

import (
	"fmt"
	"strings"
)

func main() {
	name, score := "Dent, Arthur", 87

	fmt.Println("Student scores")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(name, score)
}
