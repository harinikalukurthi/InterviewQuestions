package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range x {
		if v%2 == 0 {
			fmt.Println("Even",v)
		} else {
			fmt.Println("Odd",v)
		}
	}
}
