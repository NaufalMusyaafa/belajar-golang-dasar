package main

import "fmt"

func getFullName() (string, string) {
	return "Naufal", "Musyaafa"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)

	// jika hanya ingin mengambil salah satu nilai return value
	firstNameOnly, _ := getFullName()
	fmt.Println("First Name Only:", firstNameOnly)
}