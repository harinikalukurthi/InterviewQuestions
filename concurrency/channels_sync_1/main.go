package main

import (
	"fmt"
	"time"
)
var res=0
var value=10
func main(){
ch1:=make(chan int)
ch2:=make(chan string)
go calculateres(value,ch1)
go sendres(ch1,ch2)
<-ch2 //program wil wait it receives something into the channel2

}
func calculateres(value int,ch chan int){
	time.Sleep(3*time.Second)
	res=value*value
	ch<-res
}
func sendres(ch1 chan int,ch2 chan string){
	time.Sleep(1*time.Second)
	fmt.Println("the square of a value",<-ch1) //program wil wait it receives something into the channel1
	ch2<- "i am done with program"
}