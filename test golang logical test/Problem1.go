package main

import "fmt"

func main() {
	angka := [6]int{
		1,
		4,
		5,
		6,
		8,
		2,
	}

	fmt.Println("Random array", "\n")
	for i := 0; i < 6; i++ {
		fmt.Print(angka[i])

		for j := 0; j < angka[i]; j++ {
			fmt.Print(" |")
		}

		fmt.Println(" ")
	}

	fmt.Println("==================================", "\n", "Sorted array", "\n")

	for i := 0; i < angka[i]; i++ {
		fmt.Print(angka[i])

		for j := 0; j < angka[i]; j++ {
			fmt.Print(" |")
		}

		fmt.Println(" ")
	}

	fmt.Println("==================================", "\n", "Reverse Sorted array", "\n")

	for i := 0; i < angka[i]; i++ {
		// fmt.Print(angka[i])

		for j := 8; j >= angka[i]; j-- {
			fmt.Print(" |")
		}

		fmt.Println(" ")
	}

}
