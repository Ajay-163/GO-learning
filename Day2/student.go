package main

import "fmt"

type Student struct {
	ID     int
	Name   string
	Course string
	Marks  float64
}

func main() {
	student := Student{
		ID:     1,
		Name:   "Max",
		Course: "B.Tech CSE",
		Marks:  85,
	}

	fmt.Println("Student Details")
	fmt.Println("ID:", student.ID)
	fmt.Println("Name:", student.Name)
	fmt.Println("Course:", student.Course)
	fmt.Printf("Marks: %.2f\n", student.Marks)
}