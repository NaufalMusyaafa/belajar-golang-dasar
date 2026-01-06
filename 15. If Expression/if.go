package main

import "fmt"

func main() {
	name := "amogus"

	if name == "Naufal" {
		fmt.Println("Hello Naufal")
	} else if name == "amogus" {
		fmt.Println("Kinda sus")
	} else {
		fmt.Println("Hello Stranger")
	}

	// short statement pada if
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah pas")
	}

}