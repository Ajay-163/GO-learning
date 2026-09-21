package main

import (
	"fmt"
	"strings"
)

func main() {

	var text string

	fmt.Println("Enter a string:")
	fmt.Scanln(&text)

	// Convert string to lowercase
	text = strings.ToLower(text)

	// Count vowels
	vowels := 0

	for _, ch := range text {
		if ch == 'a' || ch == 'e' || ch == 'i' ||
			ch == 'o' || ch == 'u' {
			vowels++
		}
	}

	
	fmt.Println("Number of vowels:", vowels)
	
}