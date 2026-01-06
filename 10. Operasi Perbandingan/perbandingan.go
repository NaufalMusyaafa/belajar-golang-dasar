package main

import "fmt"

func main()  {
	var (
		nama1 string = "naufal"
		nama2 string = "naufal"
		angka1 = 10
		angka2 = 10
	)

	var result = nama1 == nama2
	var result2 = nama1 != nama2

	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(angka1 < angka2)
	fmt.Println(angka1 >= angka2)
}