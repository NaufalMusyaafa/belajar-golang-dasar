package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("perulangan ke-", counter)
	// 	counter++
	// }

	// fmt.Println("Selesai di perulangan ke-", counter)

	// for dengan init statement dan post statement

	for counter := 1; counter <= 100; counter++ {
		fmt.Println("perulangan ke-", counter)
	}

	// for range manual
	names := []string {"Muhammad", "Naufal", "Musyaafa"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for range dari slice
	for index, name := range names { // index adalah posisi ke berapa, name adalah value nya
		fmt.Println("index", index, "=", name)
	}

	// jika tanpa menggunakan index
	for _, name := range names {
		fmt.Println(name)
	}
}