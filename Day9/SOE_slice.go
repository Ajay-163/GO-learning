//Sum of elements in a slice
package main

import "fmt"

func sum(numbers []int) int {
	total := 0

	for _, value := range numbers {
		total = total + value
	}

	return total
}

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	result := sum(numbers)

	fmt.Println("Sum is:", result)
}