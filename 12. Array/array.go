package main

import "fmt"

func main() {
	var names[3] string // deklarasi string dengan panjang 3
	names[0] = "Muhammad" // index ke-0
	names[1] = "Naufal"
	names[2] = "Musyaafa"

	fmt.Println(names[2])
	fmt.Println(names[1])
	fmt.Println(names[0])

	// membuat array secara langsung
	var values = [3]int {
		50,
		60, // wajid diakhiri koma jika dienter
	}

	var value2 = [...]int { //arti titik tiga adalah jumlah arraynya tidak ingin ditentukan di awal. panjang array akan mengikuti panjang saat pertama kali dideklarasi
		50,
		60,
		70,
		80, // wajid diakhiri koma jika dienter
	}

	fmt.Println(values)
	fmt.Println(len(value2))
}