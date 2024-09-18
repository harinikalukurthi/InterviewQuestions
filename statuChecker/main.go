/*
	concurrency is that we can have multiple threads executing code, If one blockd pick the other one and worked on it.

parallelisn: multiple threads executing mutiple go routines at the exact same time*,requires multiple CPUs
*/
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	s := []string{"http://google.com", "http://ajio.com", "http://myntra.com"}
	c := make(chan string)
	for _, url := range s {
	 go check(url, c)
	}
	fmt.Println(time.Since(start))
	
	for i:=0;i<len(s);i++{
	fmt.Println(<-c)
	}
}
func check(url string, ch chan string) {

	_, err := http.Get(url)
	if err != nil {
		fmt.Println("url is down", url)
		ch <- "might be down"
		return
	}
	fmt.Printf("%s is up\n", url)
	ch <- "url is up"
}
