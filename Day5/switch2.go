//single switch case
package main
import "fmt"

func main(){
	 var day int
	fmt.Println("enter a day")
	fmt.Scanln(&day)
	switch day{
	case 1:
		fmt.Println("day 1 is monday")
	case 2:
		fmt.Println("day 2 is tuesday")
	case 3:
		fmt.Println("day 3 is wednesday")
	case 4:
		fmt.Println("day 4 is thursday")
	case 5:
		fmt.Println("day 5 is friday")
	case 6:
		fmt.Println("day 6 is saturday")
	case 7:
		fmt.Println("day 7 is sunday")
	default:
		fmt.Println("invalid day")
	}

}