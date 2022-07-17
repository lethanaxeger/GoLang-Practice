package main

import "fmt"

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "Fadlillah"
	middleName = "Bashir"
	lastName = "Hakim"

	return
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}