package main

import "fmt"

// nil adalah nilai kosong dari suatu tipe data
// nil hanya bisa digunakan pada tipe data reference seperti pointer, slice, map, channel, interface, dan function



func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{"name": name}
	}
}
func main() {
	data := NewMap("agus")
	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data)
	}
}