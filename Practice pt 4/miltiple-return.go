package main

import "fmt"

func getFullName()(string, string){
	return "Bashir", "Hakim"
}

func main(){
	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}