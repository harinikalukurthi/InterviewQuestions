package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go even(&wg)
	go odd(&wg)
	wg.Wait()

}

func even(wg *sync.WaitGroup) {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range x {
		if v%2 == 0 {
			fmt.Println("even", v)
		}
	}
	wg.Done()

}
func odd(wg *sync.WaitGroup) {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range x {
		if v%2 != 0 {
			fmt.Println("odd", v)
		}
	}
	wg.Done()

}