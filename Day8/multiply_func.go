package main
import "fmt"
 func multiply(num1 int, num2 int) int{
	product:=num1*num2
	return product
 }
 func main(){
	result:=multiply(12,20)
	fmt.Println(result)
 }