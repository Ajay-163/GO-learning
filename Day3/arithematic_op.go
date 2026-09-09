package main
import "fmt"

func main(){
	a:=15
	b:=5
	fmt.Println("using + operator", a+b)
	fmt.Println("using - operator", a-b)
	fmt.Println("using * operator", a*b)
	fmt.Println("using / operator", a/b)
	fmt.Println("using % operator", a%b)
	a++  // because go doesn't having a++ arguments in print statement
	fmt.Println("using ++ operator", a)
	b++  //because go doesn't have b-- arguments in print statement
	fmt.Println("using -- operator", b)

}