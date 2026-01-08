package main

import "fmt"

// interface kosong (any) adalah interface yang tidak memiliki method apapun
// sehingga semua tipe data otomatis mengimplementasikan interface ini
func ups() any {
	// return "Ups"
	return true
	// return valuenya bisa berupa apa saja
}

func main() {
	var kosong any = ups()
	fmt.Println(kosong)
}