package main

import "fmt"

func main() {
	var name string
	name = "amogus"
	fmt.Println(name)
	
	name = "Naufal Musyaafa"
	fmt.Println(name)

	var umur = 17
	fmt.Println(umur)

	// lebih simple lagi menggunakan := (khusus deklarasi baru)
	gender := "laki-laki"
	fmt.Println(gender)

	// deklarasi banyak variabel
	var (
		firstName = "Naufal"
		middleName = "Musyaafa"
		lastName  = "Musyaafa"
		semester = 5
	)
	fmt.Println(firstName, middleName, lastName, "berada di semester", semester)
}