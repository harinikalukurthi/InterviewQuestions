package main

import (
 "fmt"
 "time"
)

func myProcess(stopChannel chan bool) {
 for {
  select {
  case <-stopChannel:
   fmt.Println("Hey! Shantanu. Thanks for stopping my goroutine :) ")
   return
  default:
   fmt.Println("My Goroutine is running :( ")
   time.Sleep(time.Second)
  }
 }
}

func main() {
 stopChannel := make(chan bool)
 go myProcess(stopChannel)
 time.Sleep(3 * time.Second)
 stopChannel <- true
 time.Sleep(time.Second)

 fmt.Println("Main Goroutine exited")
}