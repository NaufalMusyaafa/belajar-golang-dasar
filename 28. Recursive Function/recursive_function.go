package main

import "fmt"

func factorial(n int) int {
	result := 1
	for i := n; i > 0; i-- {
		result *= i
	}
	return result
}

// recursive function adalah function yang memanggil dirinya sendiri
func factorialRecursive(value int) int {
	if value == 0 {
		return 1
	} else {
		return value * factorialRecursive(value-1) // memanggil dirinya sendiri
	}
}
func main() {
	fmt.Println(factorial(10))
	fmt.Println(factorialRecursive(5))
}