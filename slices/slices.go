package main

import "fmt"

func main() {
	// slices are similer to vector in cpp
	// with no default size and value
	var slic []int
	// with defult size
	var slic2 = make([]int, len(slic))
	// add
	slic = append(slic, 1)
	slic = append(slic, 2)
	slic = append(slic, 4)
	slic = append(slic, 5)
	copy(slic, slic2)
	fmt.Println(slic, slic2)
	// slice operator
	fmt.Println(slic[0:2])
}
