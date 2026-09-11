// menu driven calculator program

package main
import "fmt"

func main(){
	var a int
	fmt.Println("enter a value")
	fmt.Scanln(&a)
	var b int
	fmt.Println("enter b value")
	fmt.Scanln(&b)
	var mode int
	fmt.Println("enter the mode number")
	fmt.Scanln(&mode)
	switch mode{
	
	case 1:
		fmt.Println("Addition of two numbers is ",a+b)
	case 2:
		fmt.Println("Subtraction of two numbers is",a-b)
	case 3:
		fmt.Println("Multiplication of two numbers is",a*b)
	case 4:
		fmt.Println("Division of two numbers is",a/b)
	default:
		fmt.Println("Invalid mode number")
	}
	

}