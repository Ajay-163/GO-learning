// printing reverse of a number
package main 
import "fmt"
func main(){
	var n int
	fmt.Println("enter n value")
	fmt.Scanln(&n)
	rev:=0
	for n>0{
		digit:=n%10
		rev=rev*10+digit
		n=n/10
	}
	fmt.Println("reverse number is",rev)
}