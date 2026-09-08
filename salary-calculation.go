package main

import "fmt"

func main() {
	var basicSalary float64

	fmt.Print("Enter basic salary: ")
	fmt.Scanln(&basicSalary)

	da := basicSalary * 0.10
	hra := basicSalary * 0.20
	pf := basicSalary * 0.12

	grossSalary := basicSalary + da + hra
	netSalary := grossSalary - pf

	fmt.Println("\nSalary Calculation")
	fmt.Printf("Basic Salary: ₹%.2f\n", basicSalary)
	fmt.Printf("DA: ₹%.2f\n", da)
	fmt.Printf("HRA: ₹%.2f\n", hra)
	fmt.Printf("PF: ₹%.2f\n", pf)
	fmt.Printf("Gross Salary: ₹%.2f\n", grossSalary)
	fmt.Printf("Net Salary: ₹%.2f\n", netSalary)
}