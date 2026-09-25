package main

import "fmt"

type Order struct {
	orderID      int
	customerName string
	amount       float64
	products     []string
}

func main() {

	order1 := Order{
		orderID:      101,
		customerName: "Ajay",
		amount:       75000,
		products:     []string{"Laptop", "Mouse", "Keyboard"},
	}

	fmt.Println("Order ID:", order1.orderID)
	fmt.Println("Customer Name:", order1.customerName)
	fmt.Println("Amount:", order1.amount)
	fmt.Println("Products:", order1.products)
}