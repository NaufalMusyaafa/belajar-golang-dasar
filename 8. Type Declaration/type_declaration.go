package main

import "fmt"

func main() {

	type noKTP string
	
	var ktpAgus noKTP = "198237198273"

	var ktpBudi string = "298371928371"
	var ktpBudi2 noKTP = noKTP(ktpBudi)

	fmt.Println(ktpAgus)
	fmt.Println(ktpBudi)
	fmt.Println(ktpBudi2)

	

}