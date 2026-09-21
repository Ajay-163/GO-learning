package main

import (
	"fmt"
)

func main() {

	var text string

	fmt.Println("Enter a string:")
	fmt.Scanln(&text)

	reverse := ""

	for i := len(text) - 1; i >= 0; i-- {
		reverse = reverse + string(text[i])
	}

	fmt.Println("Reverse string is:", reverse)
}
