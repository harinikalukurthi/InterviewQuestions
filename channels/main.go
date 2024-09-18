package main

import (
	"fmt"
	"time"
)

func main() {
    c1 := make(chan int)
    fmt.Println("push c1: ")
	go func(){
		time.Sleep(1*time.Second)
		c1<-10
	}()
    g1 := <- c1
    fmt.Println("get g1: ", g1)
}