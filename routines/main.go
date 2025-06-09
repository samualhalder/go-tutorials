package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) { // need to pass the ref of the waitgroup
	defer wg.Done()
	fmt.Println(id)
}
func main() {
	var wg sync.WaitGroup // WaitGroup help to achive the power to wait till complition of a go routine

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go task(i, &wg)
	}
	wg.Wait()
}
