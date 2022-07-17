package main

import "fmt"

// // Traditional way to code function sayHelloWithFilter
// func sayHelloWithFilter(name string, filter func(string) string) {
// 	NameFiltered := filter(name)
// 	fmt.Println("Hello ", NameFiltered)
// }

// // Simplest way to code function sayHelloWithFilter
// // Cara ini dapat mempermudah kita bila parameter filter banyak, tinggal tambahkan parameternya di type Filter func()
// // nanti function sayHelloWothFilter akan mengikuti parameter dari type Filter func()

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func SpamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
func main() {
	sayHelloWithFilter("Eko", SpamFilter)

	filter := SpamFilter
	sayHelloWithFilter("Anjing", filter)
}
