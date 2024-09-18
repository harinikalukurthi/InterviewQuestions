package main
import (
	"sync"
)
var wg sync.WaitGroup
func main(){
	numJobs:=3
	totalJobs:=10
	ch:=make(chan int,totalJobs)
	for i:=0;i<numJobs;i++{
		go worker(i,ch)
	}




}

func worker(id int,ch <-chan int){
	
}