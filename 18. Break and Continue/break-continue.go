package main

import "fmt"

func main() {
	// break digunakan untuk menghentikan perulangan
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Berhenti di perulangan ke", i)
			break
		}
		fmt.Println("Perulangan ke", i)
	}

	// continue digunakan untuk melewati iterasi tertentu tapi melanjutkan perulangan selanjutnya
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("Perulangan ke", i)
	}
}