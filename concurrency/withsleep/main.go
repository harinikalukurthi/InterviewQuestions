package main

import (
	"fmt"
	"time"
)
func main(){
	start:=time.Now()
	fmt.Println("with concurrency")
	 go dosomething()
	 go dosomethingelse()
	 time.Sleep(2 *time.Second)
	fmt.Println("done with both methods in main")
	totaltime:=time.Since(start)
	fmt.Println(totaltime)
	}
	func dosomething(){
		time.Sleep(1*time.Second)
		fmt.Println("inside do something")
	}
	func dosomethingelse(){
		time.Sleep(1*time.Second)
		fmt.Println("inside do something else")
	}