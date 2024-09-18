//use case: to send the data from one go routine and printing the data from another go routine.
/* The Add method increases the counter by the specified amount.
The Done method decreases the counter by one. It’s typically called using defer within the goroutine to ensure it runs when the goroutine finishes.
The Wait method blocks until the counter is zero, meaning all goroutines have called Done.*/
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	x := []int{1, 2, 3}
	ch := make(chan int, len(x))
	wg.Add(2)
	go senddata(x, ch)
	go Printdata(ch)
	wg.Wait()

}

func senddata(x []int, ch chan int) {
	for i := 0; i < len(x); i++ {
		//fmt.Println("sending into channel", x[i])
		ch <- x[i]
	}
	close(ch)
	wg.Done()

}
func Printdata(ch chan int) {
	for val := range ch {
		fmt.Println(val)
	}
	wg.Done()
}