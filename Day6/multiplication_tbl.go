// Printing multiplication table
package main
import "fmt"
func main(){
	var n int
	fmt.Println("enter n value")
	fmt.Scanln(&n)
	fmt.Println("Multiplication of ",n,"table is:")
	for i:=1;i<=10;i++{
		fmt.Println(n,"*",i,"=",n*i)
	}
}