package main
import "fmt"

func main(){
	var n int
	fmt.Println("enter n value:")
	fmt.Scanln(&n)

	// even/odd
	if n%2==0{
		fmt.Println(n,"is a even number")
	}else{
		fmt.Println(n,"is a odd number")
	}
	//prime number check
	prime := true

	if n <= 1 {
		prime = false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			prime = false
			break
		}
	}

	if prime {
		fmt.Println("Prime number")
	} else {
		fmt.Println("Not a prime number")
	}
	//reverse 
	rev:=0
	temp:=n 
	for temp>0{
		digit:=temp%10
		rev=rev*10+digit
		temp=temp/10
	}
	fmt.Println("reverse number of is:",rev)
	// palindrome of number
	if n==rev{
		fmt.Println("it is a palindrome")
	}else{
		fmt.Println("it is not a palindrome")
	}
	//digits count of number
	count:=0
	tmp:=n 
	for tmp>0{
		count++
		tmp=tmp/10

	}
	fmt.Println("digit count is:",count)
}