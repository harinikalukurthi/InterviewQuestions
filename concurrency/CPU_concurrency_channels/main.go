package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan string)
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
	runtime.GOMAXPROCS(1) // no.of process that the program should use
	start := time.Now()
	go counta(ch)
	go countb(ch)
	go countc(ch)
	go countd(ch)
	go counte(ch)
	go countf(ch)
	go countg(ch)
	go counth(ch)
	for i := 0; i < 8; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("time taken is", time.Since(start))

}
func counta(ch chan string) {
	fmt.Println("A is starting")
	for i := 0; i < 1000000000; i++ {

	}
	ch <- "A is finished"

}
func countb(ch chan string) {
	fmt.Println("B is starting")
	for i := 0; i < 1000000000; i++ {

	}
	ch <- "B is finished"

}
func countc(ch chan string) {
	fmt.Println("C is starting")
	for i := 0; i < 1000000000; i++ {

	}
	ch <- "C is finished"

}
func countd(ch chan string) {
	fmt.Println("D is starting")
	for i := 0; i < 100000000; i++ {

	}
	ch <- "D is finished"

}
func counte(ch chan string) {
	fmt.Println("E is starting")
	for i := 0; i < 1000000000; i++ {

	}
	ch <- "E is finished"

}
func countf(ch chan string) {
	fmt.Println("F is starting")
	for i := 0; i < 1000000000; i++ {

	}
	ch <- "F is finished"

}
func countg(ch chan string) {
	fmt.Println("G is starting")
	for i := 0; i < 1000000000; i++ {

	}
	ch <- "G is finished"

}
func counth(ch chan string) {
	fmt.Println("H is starting")
	for i := 0; i < 1000000000; i++ {

	}
	ch <- "H is finished"

}
