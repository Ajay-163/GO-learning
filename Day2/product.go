package main

import "fmt"

type Product struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

func main() {
	product := Product{
		ID:       101,
		Name:     "Laptop",
		Quantity: 2,
		Price:    55000,
	}

	total := float64(product.Quantity) * product.Price

	fmt.Println("Product Details")
	fmt.Println("Product ID:", product.ID)
	fmt.Println("Product Name:", product.Name)
	fmt.Println("Quantity:", product.Quantity)
	fmt.Printf("Price: ₹%.2f\n", product.Price)
	fmt.Printf("Total Amount: ₹%.2f\n", total)
}