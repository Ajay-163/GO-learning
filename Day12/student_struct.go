package main

import "fmt"

type Student struct {
	name   string
	age    int
	course string
	marks  []int
}

func main() {

	student1 := Student{
		name:   "Arun",
		age:    21,
		course: "Go Programming",
		marks:  []int{80, 75, 90, 85},
	}

	fmt.Println("Name:", student1.name)
	fmt.Println("Age:", student1.age)
	fmt.Println("Course:", student1.course)
	fmt.Println("Marks:", student1.marks)
}