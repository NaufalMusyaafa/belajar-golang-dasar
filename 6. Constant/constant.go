package main

import "fmt"

func main() {
	const firstName = "Naufal"

	// firstName = "Musyaafa" // akan error karena constant tidak bisa diubah nilainya
	// const boleh tidak dipaggil
	// fmt.Println(firstName)

	// deklarasi banyak constant
	const (
		lastName  = "Musyaafa"
		age       = 17
		isMarried = false
	)
	fmt.Println(firstName, lastName)
	fmt.Println("Umur saya adalah", age)
	fmt.Println("Sudah menikah?", isMarried)
}