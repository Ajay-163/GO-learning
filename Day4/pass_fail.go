package main
import "fmt"

func main(){
	var marks int
	fmt.Println("enter the marks")
	fmt.Scanln(&marks)

	if marks>35{
		fmt.Println("the student is PASS")
	}else{
		fmt.Println("the student is FAIL")
	}
}