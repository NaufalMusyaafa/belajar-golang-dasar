package main

import "fmt"

type Customer struct {
	name, address string
	age           int
}

// membuat struct method
// struct method adalah function yang menempel di dalam struct
func (customer Customer) sayHello(name string) { //kurung sebelum nama function berisi receiver dimana receiver adalah struct yang menempel pada function tersebut dengan customer adalah nama alias dari struct Customer
	fmt.Println("Hello", name, "my name is", customer.name)
}

func main() {
	// untuk mengakses struct method, caranya sama seperti mengakses function pada umumnya
	rully := Customer{name: "Rully"}
	// sayHello() tidak bisa diakses secara langsung
	rully.sayHello("Budi")
}