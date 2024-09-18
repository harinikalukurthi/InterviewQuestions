package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)
var wg = sync.WaitGroup{}
func main(){
	runtime.GOMAXPROCS(1)
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
	wg.Add(len(links))
	for _,link:=range links{
		go checklink(link)
	}
	wg.Wait()
	fmt.Println("totaltime taken is",time.Since(start))

}
func checklink(link string){
	_,err:=http.Get(link)
		if err!=nil{
			fmt.Println("cannot reach website",link)
		}else{
		fmt.Println("website is live",link)
		}
	wg.Done()
}