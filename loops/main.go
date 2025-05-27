package main

import (
	"fmt"
)

func main() {
	// we only can use for to loop in go
	// 1. while loop
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}
	//3. infite loop
	for {
		fmt.Println("hhello")
		break
	}
	//4. classic loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// range
	// this will start from 1 & end before 3
	for i := range 3 {
		fmt.Println(i)
	}
}
