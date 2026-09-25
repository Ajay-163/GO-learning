package main

import "fmt"

// User struct
type User struct {
	name    string
	age     int
	email   string
	hobbies []string
}

// Employee struct
type Employee struct {
	name       string
	age        int
	department string
	salary     float64
	skills     []string
}

// Product struct
type Product struct {
	name     string
	price    float64
	category string
	quantity int
	tags     []string
}

// Student struct
type Student struct {
	name   string
	age    int
	course string
	marks  []int
}

// Order struct
type Order struct {
	orderID      int
	customerName string
	amount       float64
	products     []string
}

func main() {

	// Create one User object
	user1 := User{
		name:    "Ajay",
		age:     22,
		email:   "ajay@gmail.com",
		hobbies: []string{"Coding", "Reading", "Music"},
	}

	// Create one Employee object
	employee1 := Employee{
		name:       "Rahul",
		age:        25,
		department: "IT",
		salary:     45000,
		skills:     []string{"Go", "Git", "SQL"},
	}

	// Create one Product object
	product1 := Product{
		name:     "Laptop",
		price:    55000,
		category: "Electronics",
		quantity: 5,
		tags:     []string{"Computer", "Work", "Office"},
	}

	// Create one Student object
	student1 := Student{
		name:   "Arun",
		age:    21,
		course: "Go Programming",
		marks:  []int{80, 75, 90, 85},
	}

	// Create one Order object
	order1 := Order{
		orderID:      101,
		customerName: "Ajay",
		amount:       75000,
		products:     []string{"Laptop", "Mouse", "Keyboard"},
	}

	// Print User values
	fmt.Println("----- User -----")
	fmt.Println("Name:", user1.name)
	fmt.Println("Age:", user1.age)
	fmt.Println("Email:", user1.email)
	fmt.Println("Hobbies:", user1.hobbies)

	// Print Employee values
	fmt.Println("\n----- Employee -----")
	fmt.Println("Name:", employee1.name)
	fmt.Println("Age:", employee1.age)
	fmt.Println("Department:", employee1.department)
	fmt.Println("Salary:", employee1.salary)
	fmt.Println("Skills:", employee1.skills)

	// Print Product values
	fmt.Println("\n----- Product -----")
	fmt.Println("Name:", product1.name)
	fmt.Println("Price:", product1.price)
	fmt.Println("Category:", product1.category)
	fmt.Println("Quantity:", product1.quantity)
	fmt.Println("Tags:", product1.tags)

	// Print Student values
	fmt.Println("\n----- Student -----")
	fmt.Println("Name:", student1.name)
	fmt.Println("Age:", student1.age)
	fmt.Println("Course:", student1.course)
	fmt.Println("Marks:", student1.marks)

	// Print Order values
	fmt.Println("\n----- Order -----")
	fmt.Println("Order ID:", order1.orderID)
	fmt.Println("Customer Name:", order1.customerName)
	fmt.Println("Amount:", order1.amount)
	fmt.Println("Products:", order1.products)
}

