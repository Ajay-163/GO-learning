package main

import "fmt"

func minimum(numbers []int) int {
	min := numbers[0]

	for _, value := range numbers {
		if value < min {
			min = value
		}
	}

	return min
}

func main() {
	numbers := []int{10, 25, 5, 40, 15}

	result := minimum(numbers)

	fmt.Println("Minimum number is:", result)
}