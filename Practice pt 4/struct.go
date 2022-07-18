package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Malang"
	eko.Age = 18

	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	bashir := Customer{
		Name: "Bashir",
		Age: 19,
		Address: "Sawojajar",
	}
	fmt.Println(bashir)
	
}