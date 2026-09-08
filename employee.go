package main

import "fmt"

type Employee struct {			// struct  is a datatype which stores the different type of variable under single name.
    ID       int
    Name     string
    Dept     string
    Salary   float64
    JoinDate string
}

func main() {
    emp := Employee{			// assigned Employee struct to emp variable
        ID:       101,
        Name:     "maxwell",
        Dept:     "software",
        Salary:   45000,
        JoinDate: "2025-06-01",
    }

    fmt.Printf("ID:       %d\n", emp.ID)
    fmt.Printf("Name:     %s\n", emp.Name)
    fmt.Printf("Dept:     %s\n", emp.Dept)
    fmt.Printf("Salary:   ₹%.2f\n", emp.Salary)
    fmt.Printf("Join Date:%s\n", emp.JoinDate)
    
}   