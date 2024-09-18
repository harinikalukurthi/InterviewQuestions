package main

import (
	"fmt"
	"time"
)

func main() {
	const Jobs = 10
	sendJobChan := make(chan int, Jobs)
	CompleteJobChan := make(chan int, Jobs)

	for w := 1; w <= 3; w++ {
		go worker(w, sendJobChan, CompleteJobChan)
	}

for i:=1;i<=Jobs;i++{
	sendJobChan<-i
}
close(sendJobChan)
for i:=1;i<=Jobs;i++{
	<-CompleteJobChan
}

}
func worker(id int, sendchan chan int, complechan chan int) {
	for i := range sendchan {
		fmt.Println("worker", id, "Started job", i, "with", len(sendchan), "left to progress")
		time.Sleep(2*time.Second)
		fmt.Println("worker", id, "finished its job")
		complechan<-i
	}
}
