package main

import "fmt"

func main() {
	names := [...]string{
		"Muhammad",
		"Naufal",
		"Musyaafa",
		"Fadhlan",
		"Rizky",
		"Ahmad",
	}

	slice1 := names[4:6]
	fmt.Println("Slice 1:", slice1)

	slice2 := names[:3]
	fmt.Println("Slice 2:", slice2)

	slice3 := names[3:]
	fmt.Println("Slice 3:", slice3)

	slice4 := names[:]
	fmt.Println("Slice 4:", slice4)

	// cara pembuatan slice aslinya begini
	var slice5 []string = names[:] // sama seperti array, namun tanpa menentukan panjangnya
	fmt.Println("Slice 5:", slice5)

	// function pada slice

	// len(slice) => untuk mengetahui panjang slice
	panjang2 := len(slice2)
	fmt.Println("Panjang Slice 2:", panjang2)

	// cap(slice) => untuk mengetahui kapasitas slice
	kapasitas2 := cap(slice2)
	fmt.Println("Kapasitas Slice 2:", kapasitas2)

	// append(slice, value) => untuk menambahkan data pada slice
	days := [...]string {
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jum'at",
		"Sabtu",
		"Minggu",
	}
	fmt.Println("Days sebelum dirubah =", days)
	daySlice1 := days[5:] // Sabtu, Minggu
	fmt.Println("DaySlice1 =", daySlice1)
	daySlice1 [0] = "Sabtu Baru"
	daySlice1 [1] = "Minggu Baru"
	// karena daySlice1 merefer ke days, maka days juga akan berubah
	fmt.Println("Days setelah dirubah =", days)

	daySlice2 := append(daySlice1, "Libur baru") // karena kapasitas daySlice1 hanya 2, maka akan dibuat array baru
	fmt.Println("DaySlice2 =", daySlice2)

	// meskipun daySlice2 dibuat dari daySlice1, tetapi daySlice2 tidak merefer ke days. Dan karena daySlice2 melakukan append dari daySlice1 yang kapasitasnya sudah penuh, maka daySlice2 akan membuat array baru. Jadi perubahan pada daySlice2 tidak akan mempengaruhi daySlice1 maupun days
	daySlice2 [0] = "Sabtu Lama"
	fmt.Println("\n==========Hasil Setelah Mengubah nilai di daySlice2==========")
	fmt.Println("Days =", days)
	fmt.Println("DaySlice1 =", daySlice1)
	fmt.Println("DaySlice2 =", daySlice2)

	// Membuat slice dari awal menggunakan make
	var newSlice []string = make([]string, 2, 5) // 2 itu panjang slicenya, 5 itu kapasitas slicenya
	newSlice[0] = "Muhammad"
	newSlice[1] = "Muhammad"
	// newSlice[2] = "Naufal" Error, harusnya menggunakan append karena panjang slicenya cuma 2

	fmt.Println("\n==========Membuat Slice Baru Menggunakan make==========")
	fmt.Println("Slice Baru =", newSlice)
	fmt.Println("Panjang Slice baru =", len(newSlice))
	fmt.Println("Kapasitas Slice baru =", cap(newSlice))

	newSlice2 := append(newSlice, "Naufal") // tidak membuat array baru karena kapasitas masih ada
	fmt.Println("Slice Baru 2 =", newSlice2)
	fmt.Println("Panjang Slice baru 2 =", len(newSlice2))
	fmt.Println("Kapasitas Slice baru 2 =", cap(newSlice2))

	// karena tidak membuat array baru, maka ketika mengubah newSlice2, maka akan tetap merubah newSlice
	newSlice2[0] = "Ibnu"
	fmt.Println("Slice Baru 2 Setelah dirubah=", newSlice2)
	fmt.Println("Slice Baru 1 Setelah Slice 2 dirubah=", newSlice)

	// copy slice
	fmt.Println("\n==========Copy Slice==========")
	// copy(sliceTujuan, sliceSumber)
	// copy(sliceTujuan, sliceSumber) => untuk menyalin data dari slice sumber ke slice tujuan
	// jika panjang slice tujuan lebih kecil dari panjang slice sumber, maka hanya akan disalin sebanyak panjang slice tujuan
	// jika panjang slice tujuan lebih besar dari panjang slice sumber, maka hanya akan disalin sebanyak panjang slice sumber

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println("From Slice =", fromSlice)
	fmt.Println("To Slice =", toSlice)

	// perbedaan pembuatan slice dan array

	fmt.Println("\n==========Perbedaan Slice dan Array==========")
	iniArray := [...]int {1,2,3}
	iniSlice := []int {1,2,3}

	fmt.Println("Ini Array =", iniArray)
	fmt.Println("Ini Slice =", iniSlice)
}