package main

import (
	"fmt"
	"os"
	"time"
)
func main(){
ch1:=make(chan string)
ch2:=make(chan string)
ch3:=make(chan string)
 go func(){
	for{
		time.Sleep(1*time.Second)
		ch1<-"sleep for one second"
	}
}()
go func(){
for{
	time.Sleep(4*time.Second)
	ch2<-"sleep for four seconds"
}
}()
go func(){
for{
	time.Sleep(10*time.Second)
	ch3<-"exiting"
}
}()
for{
	select{
	case msg:=<-ch1:
		    fmt.Println(msg)
	case msg:=<-ch2:
		fmt.Println(msg)
	case msg:=<-ch3:
		fmt.Println(msg)
		os.Exit(0)
	}
}

}