package main

import "fmt"

func getGoodBye(name string) string {
	return "Goodbye " + name
}

func main() {
	// function as value adalah kemampuan untuk menyimpan function ke dalam sebuah variable
	goodBye := getGoodBye // bisa dilihat bahwa kita tidak menambahkan tanda kurung setelah nama function getGoodBye. Hal ini karena kita hanya ingin menyimpan functionnya ke dalam variabel saja saja, bukan memanggilnya.
	misal := getGoodBye

	// ini berarti variabel goodBye dan misal sekarang bertipe data function yang memiliki signature yang sama dengan function getGoodBye

	fmt.Println(goodBye("Naufal"))
	fmt.Println(misal("Musyaafa"))
}