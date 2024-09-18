package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
	func main() {
		for i := 0; i < 10; i++ {
			go Printroutine(i)
		}
		time.Sleep(10 *time.Second)

}

	func Printroutine(i int) {
		fmt.Println(i)
		time.Sleep(2 * time.Second)
	}
*/
var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go generate(ch)
		go printnum(ch)
	}
	wg.Wait()

}

func generate(ch chan int) {
	x := rand.Intn(10)
	ch <- x
	wg.Done()

}
func printnum(ch chan int) {
	fmt.Println(<-ch)
	wg.Done()
}
