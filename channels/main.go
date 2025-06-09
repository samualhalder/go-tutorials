package main

import (
	"fmt"
	"time"
)

func proccNum(num chan int) {
	fmt.Println(<-num)
}
func main() {
	chanNum := make(chan int)
	go proccNum(chanNum)
	chanNum <- 3
	time.Sleep(time.Second * 1)
}
