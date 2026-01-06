package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{} // membuat map kosong bernama person dengan key (di dalam kurung []) dan value bertipe string

	// person["name"] = "Muhammad Naufal"
	// person["address"] = "Indonesia"

	person := map[string]string{
		"name":    "Muhammad Naufal",
		"address": "Indonesia",
	}

	fmt.Println("value dari key 'name' di map person =", person["name"])
	fmt.Println("value dari key 'address' di map person =", person["address"])
	fmt.Println("isi dari map person =", person)
	fmt.Println("value dari key 'age' di map person =", person["age"]) // akan menghasilkan nilai default dari tipe data value, yaitu string = ""
	fmt.Println("Panjang map person =", len(person)) // mendapat jumlah data di map

	book := make(map[string]string) // deklarasi map bisa juga menggunakan make()
	book["title"] = "Belajar Golang"
	book["author"] = "Muhammad Naufal"
	book["ups"] = "Gagal"

	fmt.Println("\nIsi dari map book =", book)

	// menghapus data di map menggunakan fungsi delete
	delete(book, "ups")
	fmt.Println("Isi dari map book setelah dihapus key 'ups' =", book)
}