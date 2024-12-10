package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		c := <-ch
		fmt.Println(c)
	}()
	ch <- 1
	wg.Wait()
}
