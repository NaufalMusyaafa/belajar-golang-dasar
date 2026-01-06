package main

import "fmt"

func main() {
	name := "amogus"
	
	switch name {
	case "Naufal":
		fmt.Println("Hello Naufal")
	case "amogus":
		fmt.Println("Kinda sus")
	default:
		fmt.Println("Hello Stranger")
	}

	// short statement pada switch dengan kondisi boolean
	switch panjang := len(name); panjang > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama terlalu pendek")
	default:
		fmt.Println("Nama sudah pas")
	}

	// short statement pada switch
	switch length := len(name); {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length < 3:
		fmt.Println("Nama terlalu pendek")
	default:
		fmt.Println("Nama sudah pas")
	}
}