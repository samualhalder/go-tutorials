package main

import (
	"fmt"
	"sync"
)

type post struct {
	likes    int
	likeLock sync.Mutex
}

func (p *post) likePost(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.likeLock.Unlock()
	}()
	p.likeLock.Lock()
	p.likes += 1

}
func main() {
	post := post{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go post.likePost(&wg)
	}
	wg.Wait()
	fmt.Println(post.likes)
}
