package main
import "fmt"

func division(num1 int, num2 int) int{
	result:=num1/num2
	return result
}
func main(){
	solution:=division(100,25)
	fmt.Println(solution)
}