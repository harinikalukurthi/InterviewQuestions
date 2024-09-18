what are middlewares
what is go tooling, and name some
what are dlv package and why do you use it
what are indexing in databse and why it is important
what are context and why do we use them
what is racecondition and how to resolve it
what are methods in http package
what are mux and whay do we use it
write a go program to find the 3rd duplicate in given string
what is the difference between pod and container
which logging you are using? are you using them as part of middlewares
how exactly we have to store the credentials in go code? what precautions we have to take
what are go routinues and waitgroups?
why is that go routines do not have return parameters?
lets say you want develop a microservice from scratch, how exactly you will work. what you will consider from code to deploy. I even want to know the workspace architecture
what are function closures in the following snippet?
func main(){
	for i:=0;i<5;i++{
		wg.Add(1)
		go func(){
			fmt.Println(i)
			wg.Done()
		}()
		wg.Wait()
	}
}

what are messaging queues? how does one microservice will connect to other microservice?

we have two types of calls between microservices.
1. synchronous
2. asynchronous

synchrous calls can be made using http in spectific using rest or remote procedure calls using grpc. 
lets say A microservice send out request to B, A microservice will wait for the response from B.once it got the response from B, then it will move forward with its execution.

asynchrous calls can be made using messaging queues. 
Service A (Producer) publishes a message (task, event, or data) to the message queue.
Message Broker holds the message in a queue until Service B (Consumer) is ready to process it.
Service B (Consumer) pulls or receives the message from the queue and processes it.

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg.Add(2)
	go odd(ch1, ch2)
	go even(ch1, ch2)
	wg.Wait()

}
func odd(ch1, ch2 chan int) {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			ch1 <- i
			fmt.Println(i)
			<-ch2
		}
	}
	wg.Done()
}
func even(ch1, ch2 chan int) {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			<-ch1
			fmt.Println(i)
			ch2 <- i
		}
	}
	wg.Done()
}


write a program where we need two go routinues and wanted to print even numbers and odd numbers but in a sequence manner.
how does init functions gets triggered.
what is the difference bewteen a pod and container.
what is defer, panic?
what are goroutines?
// string-> "Hi My name is harini, people call me by the name harini.My hobbies are reading books,My hobbies are watching movies."
package main

import (
	"fmt"
	"strings"
)

func main() {
	/*ch := make(chan string)
	go print(ch)
	go print2(ch)
	time.Sleep(5 * time.Second)*/
	s := "Hi My name is harini people call me by the name harini My hobbies are reading books My hobbies are watching movies"
	removeduplicate(s)
}

/*func print(ch chan string) {
	fmt.Println("world")
	ch <- "world"
}

func print2(ch chan string) {
	<-ch
	fmt.Println("hello")
}*/

func removeduplicate(s string) {
	x := strings.Split(s, " ")
	fmt.Println(len(x))
	fmt.Println(x)
	var y []string

	m := make(map[string]int)
	for _, v := range x {
		fmt.Println(v)
		if _, ok := m[v]; !ok {
			m[v] = 1
			y=append(y,v)
		} 
	}
	fmt.Println(m)
	fmt.Println(len(m))

	fmt.Println(y)
}
//what is caching
//what is defer and write a simple program on that
//what are the things we need to look at when we want to write api calls?
// what is circular package imports and how can we resolve that
//