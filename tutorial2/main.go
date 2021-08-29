package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]*Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheCh := make(chan *Book)
	storageCh := make(chan *Book)
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10)
		wg.Add(3)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- *Book) {
			if book, ok := queryCache(id, m); ok {
				ch <- book
			}
			wg.Done()
		}(id, wg, m, cacheCh)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- *Book) {
			if book, ok := queryStorage(id); ok {
				m.Lock()
				cache[book.ID] = book
				m.Unlock()
				ch <- book
			}
			wg.Done()
		}(id, wg, m, storageCh)
		go func(i int, cacheCh, storageCh <-chan *Book, wg *sync.WaitGroup) {
			select {
			case b := <-cacheCh:
				fmt.Printf("index %v from cache %v\n", i, b)
				<-storageCh
			case b := <-storageCh:
				fmt.Printf("index %v from storage %v\n", i, b)
			}
			wg.Done()
		}(i, cacheCh, storageCh, wg)
		time.Sleep(50 * time.Millisecond)
	}
	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (*Book, bool) {
	m.RLock()
	book, ok := cache[id]
	m.RUnlock()
	return book, ok
}

func queryStorage(id int) (*Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, book := range books {
		if book.ID == id {
			return &book, true
		}
	}
	return nil, false
}
