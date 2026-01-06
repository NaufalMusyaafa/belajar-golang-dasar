package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Registration failed: Name is blacklisted.")
	} else {
		fmt.Println("User", name, "registered successfully.")
	}
}

func main() {
	// anonymus function adalah function tanpa nama yang bisa kita buat secara langsung di dalam parameter function lain.

	// ada dua cara membuat anonymus function, yaitu dengan menyimpannya ke dalam variable atau langsung membuatnya di dalam parameter function.

	// contoh menyimpan anonymus function ke dalam variable
	isBlacklisted := func(name string) bool {
		return name == "anjing"
	}
	registerUser("naufal", isBlacklisted)

	// contoh langsung membuat anonymus function di dalam parameter function
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}