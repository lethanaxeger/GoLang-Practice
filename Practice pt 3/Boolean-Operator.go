package main

import "fmt"

func main(){
	
	var nilai_ujian = 80
	var nilai_absensi = 80

	var lulusUjian = nilai_ujian > 80
	var lulusAbsensi = nilai_absensi > 80
	var lulus = lulusAbsensi && lulusUjian
	fmt.Println(lulus)
}