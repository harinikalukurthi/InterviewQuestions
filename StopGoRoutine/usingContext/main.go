// You can edit this code!
// Click here and start typing.
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ctx,cancel:=context.WithCancel(context.Background())
   // deadline:=3*time.Second
	//c1,cancel:=context.WithDeadline(context.Background(),time.Now().Add(deadline) )

	c2,cancel:=context.WithTimeout(context.Background(),3*time.Second)

	//c3:=context.WithValue(context.Background(),"a","b") // we cannot send cancellation signals through withvalue
	go processGo(c2)
	time.Sleep(time.Second * 3)
	cancel()
	fmt.Println("stopped Goroutines")

}

func processGo(c context.Context) {
	for {
		select {
		case <-c.Done():
			return
		default:
			fmt.Println("I am not yet stopped")
			time.Sleep(time.Second)
		}
	}
}
