package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	a := 42
	b := a
	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %d\n", b)

	a = 16
	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %d\n", b)

	c := &a
	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %d\n", b)
	fmt.Printf("c: %d\n", c)
	fmt.Printf("*c: %d\n", *c)

	a = 7
	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %d\n", b)
	fmt.Printf("c: %d\n", c)
	fmt.Printf("*c: %d\n", *c)
}
