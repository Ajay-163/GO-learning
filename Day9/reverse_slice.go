package main

import "fmt"

func reverse(numbers []int) []int {

	rev := []int{}

	for i := len(numbers) - 1; i >= 0; i-- {
		rev = append(rev, numbers[i])
	}

	return rev
}

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	result := reverse(numbers)

	fmt.Println("Original slice:", numbers)
	fmt.Println("Reverse slice:", result)
}