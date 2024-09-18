package main

import (
	"fmt"
	"time"
)
//var ch=make(chan string)

func main(){

	ch:=make(chan string)
	start:=time.Now()
	fmt.Println("with channels")
	go dosomething(ch)
	go dosomethingelse(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("done with both methods in main")
	totaltime:=time.Since(start)
	fmt.Println(totaltime)
	}
	func dosomething(ch chan string){
		time.Sleep(2*time.Second)
		fmt.Println("inside do something")
		ch <- "do smething channel finished"
	}
	func dosomethingelse(ch chan string){
		time.Sleep(2*time.Second)
		fmt.Println("inside do something else")
		ch <-"do smething else channel finished"
	}