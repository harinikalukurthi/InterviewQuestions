package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(6)
	fmt.Println(runtime.NumCPU())
	links := []string{
		"https://www.geeksforgeeks.org/",
		"https://stackoverflow.com/",
		"https://ajio.com",
		"https://myntra.com",
		"https://flipkart.com",
		"https://amazon.com",
	}
	ch := make(chan string,len(links))
	start := time.Now()
	for _, link := range links {
		go checklink(link, ch)
	}
	for i := 0; i < len(links); i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("totaltime taken is", time.Since(start))

}
func checklink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		ch <- "cannot reach website"
		return
	}
	ch <- "website is live"
}
