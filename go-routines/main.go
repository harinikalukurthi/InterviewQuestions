package main

import (
	"fmt"
	"sync"
)

var w int = 1000
var wg sync.WaitGroup

var mu sync.Mutex

func main() {
	fmt.Println("initail amount", w)
	wg.Add(2)

	go sell()
	go buy()
	wg.Wait()
	fmt.Println("final amount", w)
}

func sell() {

	mu.Lock()
	defer mu.Unlock()

	for i := 0; i < 500000; i++ {
		w = w + 100
	}
	wg.Done()
}
func buy() {
	mu.Lock()
	defer mu.Unlock()

	for i := 0; i < 500000; i++ {
		w = w - 100
	}
	wg.Done()
}
