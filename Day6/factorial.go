//printing factorial of a number
package main
import "fmt"
func main(){
	var n int
	fmt.Println("enter  n value")
	fmt.Scanln(&n)
	fact:=1;
	for i:=1;i<=n;i++{
		fact=fact*i
	}
	fmt.Println("factorial of the number is",fact)
}