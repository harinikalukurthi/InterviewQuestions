package main

import (
	"fmt"
	"time"
)
func main(){
	var res=0
	var value=10
	ch1:=make(chan int)
	ch2:=make(chan string)
	calculateres:=func(){
		time.Sleep(3*time.Second)
		res=value*value
		ch1<-res
	}
	sendres:=func(){
		time.Sleep(1*time.Second)
		fmt.Println("result is",<-ch1)
		ch2<-"i am done with program"
	}
	go calculateres()
	go sendres()
	<-ch2
}