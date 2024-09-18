package main

import (
	"fmt"
	"sync"
	"time"
)
func main(){
	fmt.Println("sync3")
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		var res=0
		var value=10
		ch1:=make(chan int)
		ch2:=make(chan string)
		caluculateres:=func(){
               time.Sleep(3*time.Second)
				res=value*value
				ch1<-res
		}
		sendres:=func(){
			time.Sleep(1*time.Second)
			fmt.Println("the squared value is",<-ch1)
			ch2<-"i am done with program"

		}
		go caluculateres()
		go sendres()
		<-ch2
		wg.Done()
	}()
	wg.Wait()
}