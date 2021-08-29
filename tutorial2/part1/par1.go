package part1

import (
	"fmt"
	"sync"
)

func Run(wg *sync.WaitGroup) {
	ch := make(chan string, 1)
	wg.Add(2)
	go func(ch <-chan string, wg *sync.WaitGroup) {
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	}(ch, wg)
	go func(ch chan<- string, wg *sync.WaitGroup) {
		ch <- "Hello"
		ch <- "World"
		close(ch)
		wg.Done()
	}(ch, wg)
	wg.Wait()
}
