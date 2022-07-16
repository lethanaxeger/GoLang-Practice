package main

import "fmt"

func main() {
	// counter := 1

	// for counter < 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	// For using Statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	// For Range
	slice := []string{"Fadlillah", "Bashir", "Al", "Hakim"}

	// traditional way //
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// the simplest way //
	for index, slice := range slice {
		fmt.Println("index", index, "=", slice)
	}
}
