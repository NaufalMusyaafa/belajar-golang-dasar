package main

import "fmt"

type Address struct {
	City, Province, Country string
}
func main() {
	// asterisk operator (*) adalah operator yang digunakan untuk mengakses setiap value yang direference oleh sebuah pointer
	address1 := Address{"Medan", "Sumatera Utara", "Indonesia"}
	address2 := &address1

	fmt.Println("address1 awal:", address1) 
	address2.City = "Binjai"
	fmt.Println("address2:", address2) 
	fmt.Println("address1 setelah address2 diubah:", address1) 

	address2 = &Address{"Jakarta", "DKI", "Indonesia"}
	fmt.Println("address1 setelah address2 diubah dengan reference dari struct Address:", address1) 
	fmt.Println("address2 setelah diubah:", address2)
	
	// cara agar address1 ikut berubah adalah dengan menggunakan asterisk operator (*)
	*address2 = Address{"Bandung", "Jawa Barat", "Indonesia"} // tidak perlu menggunakan & lagi karena sudah menggunakan asterisk operator
	fmt.Println("address2 setelah diubah dengan asterisk operator:", address2)
	fmt.Println("address1 setelah address2 diubah dengan asterisk operator:", address1) 

	// dengan kata lain, * pada address2 berarti setiap value yang direference oleh address2 akan diubah menjadi value yang baru sesuai dengan value pada address2. Jadi sekarang address1 akan berubah menjadi value yang sama dengan address2. Karena address1 mereference ke struct yang sama dengan address2.
}