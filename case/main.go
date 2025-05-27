package main

import "fmt"

func main() {
	age := 9
	if age > 20 {
		fmt.Println("this is an adult")
	} else if age > 13 {
		fmt.Println("this is an teenagere")
	} else {
		fmt.Println("i dont know who the fuck is this")
	}
}
