package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// new itu sama dengan &, tetapi new digunakan untuk mengembangkan pointer ke data kosong
	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.City = "Medan"
	fmt.Println("alamat1:", alamat1) // juga berubah
	fmt.Println("alamat2:", alamat2)
}