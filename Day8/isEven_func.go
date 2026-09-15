package main 
import "fmt"
func Even(num int) bool{
	if num%2==0{
		return  true
	}else{
		return false
	}
	return false
}
func main(){
	var n int
	fmt.Println("enter n value")
	fmt.Scanln(&n)
	fmt.Println(Even(n))
}