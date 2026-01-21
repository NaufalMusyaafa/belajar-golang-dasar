package main

import "fmt"

type Address struct {
	City, Province, Country string
}
func main() {
	address1 := Address{"Medan", "Sumatera Utara", "Indonesia"}
	address2 := address1

	// jika kita mengubah address2, maka address1 tidak akan berubah karena address2 adalah salinan dari address1. Ini disebut dengan pass by value

	address2.City = "Binjai"
	fmt.Println("address1:", address1) // tidak berubah
	fmt.Println("address2:", address2) // berubah menjadi binjai

	// pointer adalah kemampuan untuk membuat reference ke sebuah variable yang sama
	// dengan kemampuan pointer, kita bisa membuat pass by reference

	// untuk membuat pointer, kita bisa menggunakan operator &

	address3 := &address1

	address3.City = "Siantar"
	fmt.Println("Adress 1 setelah address 3:", address1) // berubah menjadi siantar
	fmt.Println("address3:", address3) // berubah menjadi siantar
	fmt.Println(&address3)
}