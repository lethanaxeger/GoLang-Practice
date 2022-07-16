package main

import "fmt"

func main() {

	var input = "Hakim"

	// Simple if Expression

	if input == "Bashir" {
		fmt.Println("Halo " + input)
	} else {
		fmt.Println("Maaf, siapakah gerangan?")
	}

	// If Else Expression
	if input == "Bashir" {
		fmt.Println("Halo " + input)
	} else if input == "Hakim" {
		fmt.Println("Halo " + input)
	} else if input == "Fadlil" {
		fmt.Println("Halo " + input)
	} else {
		fmt.Println("Siapakah gerangan?")
	}

	// Short Statement
	if length := len(input); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		println("Nama sudah benar")
	}
}
