package main

import "fmt"

type Student struct {
	name   string
	age    int
	course string
}

func updateStudent(student *Student) {
	student.age = 25
	student.course = "Go Programming"
}

func main() {
	student1 := Student{
		name:   "Ajay",
		age:    22,
		course: "Java",
	}

	fmt.Println("Before update:")
	fmt.Println(student1)

	updateStudent(&student1)

	fmt.Println("After update:")
	fmt.Println(student1)
}