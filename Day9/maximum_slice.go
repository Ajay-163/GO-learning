package main

import "fmt"

func maximum(numbers []int) int {
	max := numbers[0]

	for _, value := range numbers {
		if value > max {
			max = value
		}
	}

	return max
}

func main() {
	numbers := []int{10, 25, 5, 40, 15}

	result := maximum(numbers)

	fmt.Println("Maximum number is:", result)
}