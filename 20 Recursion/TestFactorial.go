package main

import "fmt"

func factorial(n int) int {
	// Stopping criteria
	if n == 1 {
		return 1
	} else {
		return factorial(n-1) * n // Recursively call yourself until the end of the return
	}
}

func main() {
	var n = 5
	var fac = factorial(n)
	fmt.Printf("The factorial of 5 is : %d", fac)
}
