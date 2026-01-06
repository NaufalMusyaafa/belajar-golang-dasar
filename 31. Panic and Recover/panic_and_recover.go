package main

import (
	"fmt"
)

// panic function adalah function yang digunakan untuk menghentikan eksekusi program secara tiba-tiba ketika terjadi error yang tidak terduga. panic biasanya digunakan ketika program mengalami kondisi yang tidak bisa diatasi, seperti akses ke index array yang diluar batas, pembagian dengan nol, atau error fatal lainnya. ketika panic dipanggil, program akan menghentikan eksekusi normalnya dan mulai mencari defer function yang ada untuk dieksekusi sebelum akhirnya program benar-benar berhenti.

// recover adalah function yang digunakan untuk menangkap nilai panic dan mencegah program dari berhenti secara tiba-tiba. recover hanya bisa dipanggil di dalam defer function. ketika recover dipanggil, program akan melanjutkan eksekusi normalnya setelah defer function selesai dieksekusi.

func endApp() {
	fmt.Println("Aplikasi selesai")

	message := recover() // recover dipanggil di dalam defer function, sehingga ketika panic terjadi, recover dapat menangkap nilai panic dan mencegah program dari berhenti secara tiba-tiba.
	if message != nil {
		fmt.Println("Terjadi error dengan message:", message)
	}
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Aplikasi error")
	} else {
		fmt.Println("Aplikasi berjalan")
	}

	// message := recover() // contoh recover yang salah karena recover harus dipanggil di dalam defer function. karena recover dipanggil di luar defer function, maka program akan tetap berhenti ketika panic dipanggil.
}
func main() {
	runApp(false)
}