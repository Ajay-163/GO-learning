package main

import "fmt"

type User struct {
	name    string
	age     int
	email   string
	hobbies []string
}

func main() {

	user1 := User{
		name:    "Ajay",
		age:     22,
		email:   "ajay@gmail.com",
		hobbies: []string{"Coding", "Reading", "Music"},
	}

	fmt.Println("Name:", user1.name)
	fmt.Println("Age:", user1.age)
	fmt.Println("Email:", user1.email)
	fmt.Println("Hobbies:", user1.hobbies)
}