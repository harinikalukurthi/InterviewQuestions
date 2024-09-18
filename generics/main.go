package main

import (
	"fmt"
)

type ordered interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~string | ~float32 | ~float64
}

// Generic function to find the maximum value
func Max[T ordered](a, b T) T {
	//func Max[T constraints.Ordered](a, b T) T {
	if a == b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Max(3, 5))       // Output: 5
	fmt.Println(Max(3.14, 2.71)) // Output: 3.14
	fmt.Println(Max("a", "b"))   // Output: b
}
