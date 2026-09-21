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

	if text == reverse {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not a palindrome")
	}
}