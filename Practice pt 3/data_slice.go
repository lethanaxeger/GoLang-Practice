package main

import "fmt"

func main(){

	fmt.Println("===================")
	fmt.Println("Normal Slice")
	fmt.Println("===================")

	var months = [...]string {
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Diubah"
	fmt.Println(slice1)

		// Append Slice Section

	fmt.Println("=====================")
	fmt.Println("Append Slice")
	fmt.Println("===================")

	var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "bashir")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

		// Make Slice Section

	fmt.Println("===================")
	fmt.Println("Make Slice")
	fmt.Println("===================")

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Bashir"
	newSlice[1] = "Hakim"
	
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

		//Copy Slide section
	
	fmt.Println("===================")
	fmt.Println("Copy Slide")
	fmt.Println("===================")

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

		// Program Array vs Slice

	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}