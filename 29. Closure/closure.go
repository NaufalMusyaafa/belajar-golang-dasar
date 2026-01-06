package main

import "fmt"

func main() {
	// closure adalah kemampuan sebuah function untuk mengakses variable di luar scope function tersebut.
	counter := 0



	// anggap saja disini adalah code yang sangat panjang.





	increment := func() {
		fmt.Println("increment")
		counter++
		// padahal tidak ada return value pada function anonymus diatas, namun kita tetap bisa mengakses variable counter yang ada di luar scope function increment tersebut. Sebaiknya gunakan closure dengan bijak agar code tetap mudah dibaca dan dipahami.
	}

	increment()
	increment()
	increment()

	fmt.Println("Counter sekarang =", counter)
}