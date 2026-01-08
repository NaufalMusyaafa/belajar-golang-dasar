package main

import "fmt"

// interface adalah tipe data abstrak yang berisi kumpulan method
type HasName interface {
	GetName() string // method GetName yang mengembalikan tipe data string
}

func SayHello(value HasName) { // function SayHello yang menerima parameter bertipe HasName
	fmt.Println("Hello", value.GetName())
}

// membuat method GetName (harus sesuai dengan kontrak di interface HasName)

// dimulai dengan membuat struct Person
type Person struct {
	name string
}

// lalu membuat method GetName untuk struct Person
func (person Person) GetName() string {
	return person.name
}

type Animal struct {
	name string
}

func (animal Animal) GetName() string {
	return animal.name
}

func main() {
	person := Person{name: "Naufal"}
	SayHello(person)
	animal := Animal{name: "Kucing"}
	SayHello(animal)
}