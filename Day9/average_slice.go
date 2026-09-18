package main

import "fmt"

func average(numbers []int) float64 {
	sum := 0

	for _, value := range numbers {
		sum = sum + value
	}

	return float64(sum) / float64(len(numbers))
}

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	result := average(numbers)

	fmt.Println("Average is:", result)
}