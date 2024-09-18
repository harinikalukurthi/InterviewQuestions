package main

import (
	"fmt"
	"sync"
	"time"
)
func main(){
	var wg sync.WaitGroup
	for i:=0;i<10;i++{
	wg.Add(1)
	go func(i int){
		var res=0
		ch1:=make(chan int)
		ch2:=make(chan string)
		caluculateres:=func(){
               time.Sleep(3*time.Second)
				res=i*i
				ch1<-res
		}
		sendres:=func(){
			time.Sleep(1*time.Second)
			fmt.Printf("the squared value of %d is %d\n",i,<-ch1)
			ch2<-"i am done with program"

		}
		go caluculateres()
		go sendres()
		<-ch2
		wg.Done()
	}(i)
	}
	wg.Wait()
}