package main

import (
	"fmt"
	"sync"
	
)

var wg = sync.WaitGroup{}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)
	wg.Add(20)
	for i := 0; i < 10; i++ {
		//fmt.Println("sum go routine", i)
		go sum(x, ch)
		//go print(ch)
	}
	for i := 0; i < 10; i++ {
		//fmt.Println("calling print go")
		go print(ch)
	}
	/*go func(ch chan int){
		for i:=0;i<10;i++{
			fmt.Println(i)
	fmt.Println(<-ch)
		}
	}(ch)*/
	wg.Wait()
	//time.Sleep(20 * time.Second)
}
func sum(x []int, ch chan int) {
	sum := 0
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	ch <- sum
	wg.Done()
}
func print(ch chan int) {
	//fmt.Println("inside print method")
	fmt.Println(<-ch)
	wg.Done()
}
