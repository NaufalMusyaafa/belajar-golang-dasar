package main

import (
	"fmt"
)

// type assertion digunakan untuk mengkonversi tipe data interface{}/any ke tipe data yang diinginkan

func random() any {
	return true
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// jangan salah mengkonversi tipe data karena akan menyebabkan panic
	// resultInt := result.(int)
	// fmt.Println(resultInt)

	// jika kita tidak yakin dengan tipe data yang dikonversi, kita bisa melakukan type assertion dengan menggunakan switch

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknown", value)
	}

}