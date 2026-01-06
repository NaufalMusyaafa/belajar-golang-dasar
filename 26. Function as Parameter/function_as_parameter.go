package main

import "fmt"

// kita juga dapat menggunakan type declaration untuk membuat tipe data function sebagai parameter function lainnya
type Filter func(string) string

func sayHelloWithFilter(name string, filter func(string) string) {
	// parameter filter adalah parameter dengan tipe data function, dengan parameter string dan return value string
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func sayHelloWithFilter2(name string, filter Filter)  {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "****"
	} else {
		return name
	}
}

func spamFilter2(name string) string {
	if name == "babi" {
		return "****"
	} else {
		return name
	}
}

func main()  {
	sayHelloWithFilter("Naufal", spamFilter) // mengirimkan function spamFilter sebagai parameter kedua. jangan lupa untuk tidak menambahkan tanda kurung setelah nama function spamFilter karena kita hanya ingin mengirimkan functionnya saja, bukan memanggilnya.

	// kita juga bisa menyimpan function spamFilter ke dalam variable terlebih dahulu baru mengirimkannya sebagai parameter
	filter := spamFilter
	sayHelloWithFilter("anjing", filter)

	sayHelloWithFilter2("babi", spamFilter2)

}

