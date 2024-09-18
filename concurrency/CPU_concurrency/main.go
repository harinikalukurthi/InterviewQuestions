package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)
var wg=sync.WaitGroup{}
func main(){
fmt.Println(runtime.GOMAXPROCS(0))
runtime.GOMAXPROCS(16) // no.of process that the program should use
wg.Add(8)
start:=time.Now()
go counta()
go countb()
go countc()
go countd()
go counte()
go countf()
go countg()
go counth()
wg.Wait()
fmt.Println("time taken is",time.Since(start))

}
func counta(){
	fmt.Println("A is starting")
	for i:=0;i<1000000000;i++{

	}
	fmt.Println("A is finished")
	wg.Done()
}
func countb(){
	fmt.Println("B is starting")
	for i:=0;i<1000000000;i++{
		
	}
	fmt.Println("B is finished")
	wg.Done()
}
func countc(){
	fmt.Println("C is starting")
	for i:=0;i<1000000000;i++{
		
	}
	fmt.Println("C is finished")
	wg.Done()
}
func countd(){
	fmt.Println("D is starting")
	for i:=0;i<100000000;i++{
		
	}
	fmt.Println("D is finished")
	wg.Done()
}
func counte(){
	fmt.Println("E is starting")
	for i:=0;i<1000000000;i++{
		
	}
	fmt.Println("E is finished")
	wg.Done()
}
func countf(){
	fmt.Println("F is starting")
	for i:=0;i<1000000000;i++{
		
	}
	fmt.Println("F is finished")
	wg.Done()
}
func countg(){
	fmt.Println("G is starting")
	for i:=0;i<1000000000;i++{
		
	}
	fmt.Println("G is finished")
	wg.Done()
}
func counth(){
	fmt.Println("H is starting")
	for i:=0;i<1000000000;i++{
		
	}
	fmt.Println("H is finished")
	wg.Done()
}