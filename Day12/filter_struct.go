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

	var department string

	fmt.Print("Enter department: ")
	fmt.Scanln(&department)

	fmt.Println("Employees in", department, ":")

	for _, employee := range employees {

		if employee.department == department {
			fmt.Println("Name:", employee.name)
			fmt.Println("Age:", employee.age)
			fmt.Println("Salary:", employee.salary)
			fmt.Println()
		}
	}
}