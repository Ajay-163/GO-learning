//printing reverse of a string
package main
import "fmt"
func main(){
	var str string
	fmt.Println("enter string ")
	fmt.Scanln(&str)
	rev:=""
	for i:=len(str)-1;i>=0;i--{
		rev=rev+string(str[i])
	}
	fmt.Println(rev)
}