package main

import (
	"fmt"
	"runtime"
	"time"
)
func main(){
fmt.Println(runtime.GOMAXPROCS(0))
runtime.GOMAXPROCS(1)
start:=time.Now()
counta()
countb()
countc()
countd()
counte()
countf()
countg()
counth()

fmt.Println("time taken is",time.Since(start))

}
func counta(){
	fmt.Println("A is starting")
	for i:=0;i<1000000000;i++{

	}
	fmt.Println("A is finished")
}
func countb(){
	fmt.Println("B is starting")
	for i:=0;i<1000000000;i++{
		
	}
	fmt.Println("B is finished")
}
func countc(){
	fmt.Println("C is starting")
	for i:=0;i<1000000000;i++{
		
	}
	fmt.Println("C is finished")
}
func countd(){
	fmt.Println("D is starting")
	for i:=0;i<100000000;i++{
		
	}
	fmt.Println("D is finished")
}
func counte(){
	fmt.Println("E is starting")
	for i:=0;i<1000000000;i++{
		
	}
	fmt.Println("E is finished")
}
func countf(){
	fmt.Println("F is starting")
	for i:=0;i<1000000000;i++{
		
	}
	fmt.Println("F is finished")
}
func countg(){
	fmt.Println("G is starting")
	for i:=0;i<1000000000;i++{
		
	}
	fmt.Println("G is finished")
}
func counth(){
	fmt.Println("H is starting")
	for i:=0;i<1000000000;i++{
		
	}
	fmt.Println("H is finished")
}