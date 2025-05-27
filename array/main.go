package main

import "fmt"

func main() {
	var arr [4]int
	arr[0] = 2
	arr[1] = 3
	arr[2] = 5
	fmt.Println(arr)

	// loop in array

	for i := range len(arr) {
		fmt.Println(arr[i])
	}
	// diract value assign

	newArr := [2]int{1, 2}
	new2DArr := [2][2]int{{1, 2}, {3, 9}}

	fmt.Println(newArr, new2DArr)
}
