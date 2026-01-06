package main

import "fmt"

// defer adalah sebuah statement yang digunakan untuk menunda eksekusi sebuah function hingga function yang memanggilnya selesai dieksekusi. defer function akan tetap dieksekusi meskipun terjadi error/panic di dalam function yang memanggilnya.

func logging()  {
	fmt.Println("Selesai memanggil function")
}

func application()  {
	defer logging() // function logging akan dieksekusi setelah function application selesai dieksekusi
	fmt.Println("Aplikasi berjalan")
	
}

func main() {
	application()
}