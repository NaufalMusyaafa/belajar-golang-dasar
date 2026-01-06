package main

import "fmt"

func main() {
	var ujian = 10
	var tugas = 90

	var lulusUjian = ujian >= 85
	var lulusTugas = tugas >= 85

	var lulus = lulusTugas || lulusUjian

	fmt.Println("apakah dia lulus=", lulus)
}