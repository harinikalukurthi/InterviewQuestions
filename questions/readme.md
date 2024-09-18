1. what are middlewares
2. what is go tooling, and name some
3. what are dlv package and why do you use it
4. what are indexing in databse and why it is important
5. what are context and why do we use them
6. what is racecondition and how to resolve it
7. what are methods in http package
8. what are mux and whay do we use it
9. write a go program to find the 3rd duplicate in given string
10. what is the difference between pod and container
11. which logging you are using? are you using them as part of middlewares
12. how exactly we have to store the credentials in go code? what precautions we have to take
13. what are go routinues and waitgroups?
14. why is that go routines do not have return parameters?
15. lets say you want develop a microservice from scratch, how exactly you will work. what you will consider from code to deploy. I even want to know the workspace architecture
16. what are function closures in the following snippet?
```
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
```

17. what are messaging queues? how does one microservice will connect to other microservice?
```
we have two types of calls between microservices.
1. synchronous
2. asynchronous

synchrous calls can be made using http in spectific using rest or remote procedure calls using grpc. 
lets say A microservice send out request to B, A microservice will wait for the response from B.once it got the response from B, then it will move forward with its execution.

asynchrous calls can be made using messaging queues. 
Service A (Producer) publishes a message (task, event, or data) to the message queue.
Message Broker holds the message in a queue until Service B (Consumer) is ready to process it.
Service B (Consumer) pulls or receives the message from the queue and processes it.
```

18. write a program where we need two go routinues and wanted to print even numbers and odd numbers but in a sequence manner.
```
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
```

19. how does init functions gets triggered.
20. what is the difference bewteen a pod and container.
21. what is defer, panic?
22. what are goroutines?
23. use single channel to connect both goroutines. I want to send the data from one channel and want to receive the channel.
24. what is goroutines leak
25. what is that we have to look while writing a testcase for a function in go
26. what is hashing and how can we do that.
27. write a api call to change the passeword of a user.
```
package main
m:=make(map[string]string)


func main(){
 
r:=gin.Default()
r.PUT(“/:change”,changePass)
r.Run()
}



func changePass(r *gin.Context){
newpass:=r.Param("change")
If pass,ok:=m[user-name];ok{
If pass!=newpass{
m[user-name]=newpass
}else{
r.JSON(statuscode,”string”)
}
}

}

 Curl -u “user”: “password” -X PUT https://localhost:<port/<end-point>
```
28. What are different types of authentication?
Basic auth
JWT
Oidc
Saml
Bearer token
29. Write a SQL query to list the emp name and his/her manager’s name.
Employee: [Name, ID, MgrID] where Name is the employee name and ID is the identity of an employee.
30. Normalize the table 
```
Restaurant : [Name, Photo, FoodName, FoodPhoto, Price]
```
31. get the sizes of the following
```
type myStruct1 struct {
myBool bool // 1 byte
myFloat float64 // 8 bytes
myInt int32 // 4 bytes
}
type myStruct2 struct {
myFloat float64 // 8 bytes
myBool bool // 1 byte
myInt int32 // 4 bytes
}
```

solution: the first struct would be 24 and second  would be 16 bytes

32. please write a program to give the second largest number from an array with variable length of arrays

```
func main() {
	getsecond(3, 2, 4, 5, 9, 10)
	getsecond(0, 4, 1, 2)

}
func getsecond(a ...int) {
	sort.Ints(a)
	fmt.Println(a[len(a)-2])
}

```

33. write a program explaining panic, defer and recover.
```
package main

import (
	"fmt"
	"sort"
)

func main() {
	res:=divide(2,1)
	fmt.Println(res)
}

func divide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic")
		}
	}()
	if b == 0 {
		panic("cannot divide by zero")
	}
	return a / b
}
```

34. how can we channelize the synchronisation? Hint: use channels to make the main go routine to wait.
35. what is the use of init function in golang
