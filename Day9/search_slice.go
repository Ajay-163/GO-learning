//search for an element in a slice
package main

import "fmt"

func search(numbers []int, target int) bool {

	for _, value := range numbers {
		if value == target {
			return true
		}
	}

	return false
}

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	var target int

	fmt.Println("Enter number to search:")
	fmt.Scanln(&target)

	if search(numbers, target) {
		fmt.Println("Number found")
	} else {
		fmt.Println("Number not found")
	}
}
