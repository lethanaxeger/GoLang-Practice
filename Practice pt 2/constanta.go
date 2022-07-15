package main

import "fmt"

func main(){
	const (
		firstName = "Eko"
		lastName = "main"
	)
	fmt.Println(firstName + ", " + lastName)
	cons()
	// pasti error jika

	// firstName = "Tidak Bisa Diubah"
	// lastName = "Tidak Bisa Diubah"

	
}

func cons(){
	const firstName = "Siapa"
	const lastName = "cons"

	fmt.Println(firstName + ", " + lastName)
}