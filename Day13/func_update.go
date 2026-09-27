package main

import "fmt"

func update(x int) {
	x = 100
	fmt.Println("Inside function:", x)
}

func main() {

	x := 10

	fmt.Println("Before:", x)

	update(x)

	fmt.Println("After:", x)
}