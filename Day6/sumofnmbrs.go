//printing sum of n  numbers:
package main
import "fmt"
func main(){
	var n int
	sum:=0;
	fmt.Println("enter n values")
	fmt.Scanln(&n)
	for i:=1;i<=n;i++{
		sum=sum+i
	}
	fmt.Println("sum of numbers is :",sum)
}