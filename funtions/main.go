package main

import "fmt"

//funtion can return multiple vairbale in go
func fun() (int, string, bool, int, string) {
	return 2, "hello", true, 2, "hey"
}

// in go funcitons are first class citizens ie you can take/returnt or store a fucntion in side a variabel
func newFun(fn func(a int) int) int {
	return fn(2)
}

func main() {
	var ans [2]int
	ans[0] = 1
	a := func(val int) int {
		return val * 2
	}
	fmt.Println(newFun(a))
}
