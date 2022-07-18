package main

import "fmt"

func logging(){
	fmt.Println("Selesai memanggil function")
}

func runAppAplication(value int) {

	// // Defer harus diletakkan diatas
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
	// logging()
}

func main(){
	runAppAplication(0)
}