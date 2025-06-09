package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// file, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(file.Name())
	buf := make([]byte, 11)
	l, err := f.Read(buf)
	if err != nil {
		panic(err)
		fmt.Println(l)
	}
	for i := range buf {
		fmt.Print(string(buf[i]))
	}
}
