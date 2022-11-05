package main

import (
	"concurrency/book"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]book.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheCh := make(chan book.Book)
	dbCh := make(chan book.Book)
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- book.Book) {
			if b, ok := queryCache(id, m); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, cacheCh)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- book.Book) {
			if b, ok := queryDatabase(id, m); ok {
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, wg, m, dbCh)
		go func(chacheCh, dbCh <-chan book.Book) {
			select {
			case b := <-chacheCh:
				fmt.Println("from chache")
				fmt.Println(b)
				<-dbCh
			case b := <-dbCh:
				fmt.Println("from database")
				fmt.Println(b)
			}

		}(cacheCh, dbCh)
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("book not found with id")
	}
	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (book.Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (book.Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range book.Books {
		if b.ID == id {
			return b, true
		}
	}
	return book.Book{}, false
}
