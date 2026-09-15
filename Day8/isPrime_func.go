package main

import "fmt"

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }

    for i := 2; i < n; i++ {
        if n%i == 0 {
            return false
        }
    }

    return true
}

func main() {
    var n int

    fmt.Println("Enter a number:")
    fmt.Scanln(&n)

    if isPrime(n) {
        fmt.Println("Prime number")
    } else {
        fmt.Println("Not a prime number")
    }
}