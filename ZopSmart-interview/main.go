package main

import (
	"fmt"
	"time"
)

func main() {

	chA := make(chan string)
	chB := make(chan string)
	go callA(chA, chB)
	go callB(chA, chB)

	time.Sleep(1 * time.Second)
	chA <- "AAA"
	time.Sleep(1 *time.Second)
}
func callA(ca, cb chan string) {
	for {
		select {
		case msg:=<-ca:
			cb <- " sent to B"
			fmt.Println(msg)
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("no data sent to channel A")
		}
	}

}
func callB(ca, cb chan string) {
	for {
		select {
		case msg:=<-cb:
			ca <- "sent to A"
			fmt.Println(msg)
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("no data sent to channel B")
		}
	}

}
