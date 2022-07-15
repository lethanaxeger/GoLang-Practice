package main

import "fmt"

func main() {
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println("Konversi pertama")
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)
	fmt.Println("===========================")

	positif_semua()
	konversi_dua()
}

func positif_semua() {
	var nilai32 int32 = 127
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)
	fmt.Println("================")
}

func konversi_dua() {
	var name = "Bashir"
	var e = name[0]
	var eString = (e)

	fmt.Println("Konversi ke 2")
	fmt.Println(name)
	fmt.Println(eString)
}
