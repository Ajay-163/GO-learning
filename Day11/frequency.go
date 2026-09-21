// characters frequency
package main

import "fmt"

func main() {

	var text string

	fmt.Println("Enter a string:")
	fmt.Scanln(&text)

	frequency := make(map[byte]int)

	for i := 0; i < len(text); i++ {
		frequency[text[i]]++
	}

	for ch, count := range frequency {
		fmt.Println(string(ch), "=", count)
	}
}