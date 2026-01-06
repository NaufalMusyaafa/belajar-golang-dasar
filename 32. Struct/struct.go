package main

import "fmt"

// struct adalah sebuah template data yang digunakan untuk menggabungkan nolai atau lebih tipe data ke dalam satu kesatuan
// data di struct disebut field
// sederhananya struct adalah kumpulan dari field yang memiliki tipe data berbeda

type Person struct {
	Name, Address string
	Age           int
}

// kita tidak bisa langsung menggunakan struct, tetapi kita bisa membuat data/object dari struct tersebut

func main() {
	var person1 Person

	person1.Name = "Naufal"
	person1.Address = "Jl. Merdeka No. 1"
	person1.Age = 20

	fmt.Println(person1)
	fmt.Println(person1.Name)
	fmt.Println(person1.Address)
	fmt.Println(person1.Age)

	// cara lain membuat object dari struct
	Budi := Person{
		Name:   "Budi",
		Age:    25,
		Address: "Jl. Sudirman No. 2",
	}
	fmt.Println(Budi)

	// cara lain membuat object dari struct tanpa menyebutkan nama field
	Susi := Person{"Susi", "Jl. Thamrin No. 3", 22}
	fmt.Println(Susi)

}