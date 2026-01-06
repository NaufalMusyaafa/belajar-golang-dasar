package main

import "fmt"

func getHello(name string) string { // function dengan return value string
	return "Hello, " + name + "!"
}

func main() {
	// memanggil function dan menampung return value nya ke dalam variable
	result := getHello("Naufal")
	fmt.Println(result)
}