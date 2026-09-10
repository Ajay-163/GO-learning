package main
import "fmt"

func main(){
	var n int
	fmt.Println("enter a number:")
	fmt.Scanln(&n)

	if n>0{
		fmt.Println("it is a postive number")
	} else if n==0{
		fmt.Println("it is a whole number")
	}else{
		fmt.Println("it is a negtive number")
	}
}