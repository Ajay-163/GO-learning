package main

import "fmt"

type Employee struct {
	name       string
	age        int
	department string
	salary     float64
	skills     []string
}

func main() {

	employee1 := Employee{
		name:       "Rahul",
		age:        25,
		department: "IT",
		salary:     45000,
		skills:     []string{"Go", "Git", "SQL"},
	}

	fmt.Println("Name:", employee1.name)
	fmt.Println("Age:", employee1.age)
	fmt.Println("Department:", employee1.department)
	fmt.Println("Salary:", employee1.salary)
	fmt.Println("Skills:", employee1.skills)
}