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

	// Switch Expression
	name := "Fadlillah"

	switch name {
	case "Fadlillah":
		fmt.Println("Hello Fadlillah")

	case "Bashir":
		fmt.Println("Hello Bashir")

	case "Hakim":
		fmt.Println("Hello Hakim")

	default:
		fmt.Println("Siapakah gerangan?")
	}

	// Short Statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")

	case false:
		fmt.Println("Nama sudah benar")

	}

	// Switch without any condition
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")

	case length > 5:
		fmt.Println("Nama lumayan panjang")

	default:
		fmt.Println("Nama sudah benar")
	}
}
