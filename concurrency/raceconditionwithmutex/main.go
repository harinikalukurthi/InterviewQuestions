package main

import (
	"fmt"
	"runtime"
	"sync"
)
var amount=1000
var wg sync.WaitGroup
var mu sync.Mutex
func main(){
	runtime.GOMAXPROCS(8)
	fmt.Println("the amount initailly", amount)
	wg.Add(2)
	go sell()
	go buy()
	wg.Wait()
	fmt.Println("the final amount is",amount)
}
func sell(){
	for i:=0;i<30000;i++{
		mu.Lock()
		amount-=100
		mu.Unlock()
		//fmt.Println(amount)
	}
	wg.Done()
}
func buy(){
	for i:=0;i<30000;i++{
		mu.Lock()
		amount+=100
		mu.Unlock()
		//fmt.Println(amount)
	}
	wg.Done()
}