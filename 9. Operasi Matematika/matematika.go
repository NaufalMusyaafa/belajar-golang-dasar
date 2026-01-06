package main

import "fmt"

func main() {
	var (
		a = 10
		b = 20
		d = 2
		c = a + b * (d % a)
	)
	
	fmt.Println("c =", c)
	
	// augmented assignment
	var i = 10
	fmt.Println("i =", i)
	i += 10 // i = i + 10
	fmt.Println("i += 10 =", i,)

	// unary operator

	var u = 10
	fmt.Println("u =",u)
	u++ // u = u + 1
	fmt.Println("u++ =",u)
}