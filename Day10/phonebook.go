package main

import "fmt"

func main() {

	phoneBook := make(map[string]string)

	// Add contacts
	phoneBook["Ajay"] = "9876543210"
	phoneBook["Ravi"] = "9876501234"
	phoneBook["Rahul"] = "9123456780"

	// Search contact
	var name string

	fmt.Println("Enter name to search:")
	fmt.Scanln(&name)

	phone, exists := phoneBook[name]

	if exists {
		fmt.Println("Phone number is:", phone)
	} else {
		fmt.Println("Contact not found")
	}
}