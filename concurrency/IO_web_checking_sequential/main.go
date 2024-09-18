package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)
func main(){
	runtime.GOMAXPROCS(16)
	fmt.Println(runtime.NumCPU())
	links:=[]string{
		"https://www.geeksforgeeks.org/",
		"https://stackoverflow.com/",
		"https://ajio.com",
		"https://myntra.com",
		//"https://flipkart.com",
		"https://amazon.com",
	}
	start := time.Now()
	for _,link:=range links{
		checklink(link)
	}
	fmt.Println("totaltime taken is",time.Since(start))
}
func checklink(link string){
	_,err:=http.Get(link)
		if err!=nil{
			fmt.Println("cannot reach website",link)
			return
		}
		fmt.Println("website is live",link)
}