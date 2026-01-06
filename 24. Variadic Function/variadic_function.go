package main

import "fmt"

// variadic function adalah function yang bisa menerima parameter dalam jumlah yang tidak terbatas. Parameter dari variadic fuction harus berada di posisi paling akhir dari parameter function tersebut.
func sumAll(numbers ...int) int { // parameter numbers merupakan variadic parameter
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// jika menggunakan slice secara manual
func sumAllSlice(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(sumAll(10,10,10,10))
	fmt.Println(sumAllSlice([]int{10,10,10})) // jika menggunakan slice secara manual

	// jika ingin mengirimkan slice ke variadic function, kita bisa menambahkan ... setelah nama variadic parameter
	slice := []int{20,20,20,20}
	fmt.Println(sumAll(slice...))
}

