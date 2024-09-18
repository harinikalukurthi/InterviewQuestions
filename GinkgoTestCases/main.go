package main

import "fmt"

func main() {
	x := 2
	y := 4
	fmt.Println(Sum(x, y))

}
func Sum(a, b int) int {
	return a + b
}
