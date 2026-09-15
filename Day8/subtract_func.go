package main
import "fmt"

func subtract(num1 int, num2 int) int{
	result:=num1-num2
	return result
}
func main(){
	fmt.Println(subtract(50,20))
}