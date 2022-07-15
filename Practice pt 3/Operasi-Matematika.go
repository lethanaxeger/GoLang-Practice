package main

import "fmt"

func Augmented_Assignments(){
	var i = 10
	i += 10 // i = i + 10

	fmt.Println("Augmented Assignments")
	fmt.Println(i)

}

func Unary_Operator(){
	var i = 10
	i += 10
	i++
	
	fmt.Println("Unary Operation")
	fmt.Println(i)
}

func main(){
	Augmented_Assignments()
	Unary_Operator()
}