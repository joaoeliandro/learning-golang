package main

import "fmt"

func main() {
	x := 10
	y := "Hello World!"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	x = 20
	fmt.Printf("x: %v, %T\n", x, x)
}
