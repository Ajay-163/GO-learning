package main
import "fmt"

func main(){
	var month int
	fmt.Println("enter the month number")
	fmt.Scanln(&month)
	switch month{
	case 1:
		fmt.Println("month is January")
	case 2:
		fmt.Println("month is february")
	case 3:
		fmt.Println("month is March")
	case 4:
		fmt.Println("month is April")
	case 5:
		fmt.Println("month is May")
	case 6:
		fmt.Println("month is June")
	case 7:
		fmt.Println("month is July")
	case 8:
		fmt.Println("month is August")
	case 9:
		fmt.Println("month is September")
	case 10:
		fmt.Println("month is October")
	case 11:
		fmt.Println("month is November")
	case 12:
		fmt.Println("month is December")
	default:
		fmt.Println("invalid month number")
		
	}
}