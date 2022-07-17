package main

import "fmt"

func main(){

	// break
	for i := 0; i < 10; i++ {
		fmt.Println("Perulangan ke ", i)
		if i == 5 {
			break
		}
	}
	fmt.Println("=============================")

	// continue
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println("perulangan ke", i)
	}
}