package main

import "fmt"
func main(){

	n:=4
	fmt.Println(fib(n))

}
func fib(n int)int{
	if n<=1{
		return n
	}
		return fib(n-2)+fib(n-1)
}