package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "John"
	middleName = "Doe"
	// lastName = "Smith" // secara default bernilai string kosong
	return
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println("First Name:", a)
	fmt.Println("Middle Name:", b)
	fmt.Println("Last Name:", c)
}