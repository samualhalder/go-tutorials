package main

import "fmt"

func main() {
	mp := make(map[string]int)
	mp["samual"] = 1
	mp["samu"] = 10
	mp["samuaa"] = 11
	// delte
	delete(mp, "samu")
	//clear
	clear(mp)
	fmt.Println(mp["samu"])
}
