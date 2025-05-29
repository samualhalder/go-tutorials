package main

import "fmt"

func fn1(a *int) {
	fmt.Println(*a)
}

func main() {
	a := 2
	fn1(&a)
}
