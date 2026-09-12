//printing even/odd numbers
package main
import "fmt"
func main(){
	n:=10
	for i:=1;i<=n;i++{
		if i%2==0{
			fmt.Println("i is even number",i)
		}else{
			fmt.Println("i is odd number",i)
		}
	}
}