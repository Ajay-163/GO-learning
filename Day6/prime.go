//printing prime numbers from 1 to 100
package main
import "fmt"
func main(){
	
	// Iterate from 1 to 100
	for num := 1; num <= 100; num++ {
		// Assume the number is prime
		isPrime := true

		// Handle 1 specifically as it is not prime
		if num == 1 {
			isPrime = false
		} else {
			// Checking for divisors from 2 to num-1
			//  Check up to num/2 is sufficient
			for i := 2; i <= num/2; i++ {
				if num%i == 0 {
					isPrime = false
					break // Not a prime, exit inner loop
				}
			}
		}

		// Print the number if it is prime
		if isPrime {
			fmt.Println(num)
		}
	}
}