package main
import "fmt"

func main(){
	var age int
	fmt.Print("enter age of person:")
	fmt.Scanln(&age)

	if age>18{
		fmt.Println("Person is  major")
	} else{
		fmt.Println("Person is Minor")
	}

}