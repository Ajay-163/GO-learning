package main

import "fmt"

type Employee struct {
	name       string
	age        int
	department string
	salary     float64
}

func main() {

	employees := []Employee{
		{
			name:       "Ajay",
			age:        22,
			department: "IT",
			salary:     40000,
		},
		{
			name:       "Rahul",
			age:        25,
			department: "HR",
			salary:     45000,
		},
		{
			name:       "Arun",
			age:        28,
			department: "IT",
			salary:     55000,
		},
	}

	var searchName string

	fmt.Print("Enter employee name: ")
	fmt.Scanln(&searchName)

	for _, employee := range employees {

		if employee.name == searchName {
			fmt.Println("Employee found!")
			fmt.Println("Name:", employee.name)
			fmt.Println("Age:", employee.age)
			fmt.Println("Department:", employee.department)
			fmt.Println("Salary:", employee.salary)
			return
		}
	}

	fmt.Println("Employee not found")
}