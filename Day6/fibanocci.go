//printing fibanocci series of a number
package main
import "fmt"
func main(){
	var n int
	fmt.Println("enter a number")
	fmt.Scanln(&n)
	a:=0
	b:=1
	for i:=1;i<n;i++{
		fmt.Print(a," ")
		next:=a+b 
		a=b 
		b=next;
	}
}