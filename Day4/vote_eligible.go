package main
import "fmt"

func main(){
	var age int
	fmt.Println("enter the age of candidate")
	fmt.Scanln(&age)
	 
	if age>=18{
		fmt.Println("Candidate is eligible for vote")
	}else{
		
	fmt.Println("candidate is not eligible for vote")
	}
}