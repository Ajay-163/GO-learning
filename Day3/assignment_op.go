package main
import "fmt"

func main(){
	a:=10
	b:=5
	c:=a
	fmt.Println("simple assignment = ",c)
	c+=b
	fmt.Println("addition assignment +=",c)
	c-=a
	fmt.Println("subtraction assignment -=",c)
}