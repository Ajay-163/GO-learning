package main

import "fmt"

type Product struct {
	name     string
	price    float64
	category string
	quantity int
	tags     []string
}

func main() {

	product1 := Product{
		name:     "Laptop",
		price:    55000,
		category: "Electronics",
		quantity: 5,
		tags:     []string{"Computer", "Work", "Office"},
	}

	fmt.Println("Name:", product1.name)
	fmt.Println("Price:", product1.price)
	fmt.Println("Category:", product1.category)
	fmt.Println("Quantity:", product1.quantity)
	fmt.Println("Tags:", product1.tags)
}