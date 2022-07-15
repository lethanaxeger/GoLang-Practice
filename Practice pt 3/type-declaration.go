package main

import "fmt"

func main(){
	type NoKTP string
	type Kerja bool
	
	var noKtpBashir NoKTP = "350711223344"
	var KerjaStatus Kerja = true
	fmt.Println(noKtpBashir)
	fmt.Println(KerjaStatus)
}