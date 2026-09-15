package main
import "fmt"

func factorial(num int) int{
	fact:=1
	for i:=1;i<=num;i++{
		fact=fact*i
	}
	return fact
	
}
func main(){
	var n int
	fmt.Println("enter n value")
	fmt.Scanln(&n)
	result:=factorial(n)
	fmt.Println(result)
}