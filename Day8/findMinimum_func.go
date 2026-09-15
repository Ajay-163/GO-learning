package main 
import "fmt"
func minimum(num1 int,num2 int,num3 int) int{
	if num1<num2 && num1<num3{
		return num1
	}else if num2<num1 && num2<num3{
		return num2
	}else{
		return num3
	}
	
}
func main(){
	var n1,n2,n3 int
	fmt.Println("enter n1,n2,n3 values")
	fmt.Scanln(&n1,&n2,&n3)
	fmt.Println(minimum(n1,n2,n3))
}