package main

import "fmt"

var y int

func main() {
	x := 10
	y := "Hello World!"
	result, _ := fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	x = 20
	fmt.Printf("x: %v, %T\n", x, x)

	test_func(result)
}

func test_func(result int) {
	fmt.Printf("Result: %v\n", result)

	fmt.Println(y)
}
