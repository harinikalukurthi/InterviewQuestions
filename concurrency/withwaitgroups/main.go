package main

import (
	"fmt"
	"sync"
	"time"
)
var wg=sync.WaitGroup{}
func main(){

	wg.Add(2)
	start:=time.Now()
	fmt.Println("with waitgroups")
	go dosomething()
	go dosomethingelse()
    wg.Wait()
	 fmt.Println("done with both methods in main")
	totaltime:=time.Since(start)
	fmt.Println(totaltime)
	}
	func dosomething(){
		time.Sleep(2*time.Second)
		fmt.Println("inside do something")
		wg.Done()
	}
	func dosomethingelse(){
		time.Sleep(2*time.Second)
		fmt.Println("inside do something else")
		wg.Done()
	}