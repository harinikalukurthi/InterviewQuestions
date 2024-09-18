package main

import "fmt"

var s []int

func main() {
Push(2)
fmt.Println(s)
Push(3)
Display()
Push(4)
fmt.Println(s)
fmt.Println(pop())
fmt.Println(s)
Display()
}

func Push(x int) {
	s = append(s, x)

}
func pop() (x int) {
	//x=s[len(s)-1]
   //s=(s[:len(s)-1])  stack 

   x=s[0]
   s=s[1:]  //queue
   return x
}
func Display() {
	fmt.Println(s)
}
