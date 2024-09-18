get the sizes of the following

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

sol: the first struct would be 24 and second  would be 16 bytes


please write a program to give the second largest number from an array with variable length of arrays



write a program explaining panic, defer and recover.

how can we channelize the synchronisation


solutions:


package main

import (
	"fmt"
	"sort"
)

func main() {
	getsecond(3, 2, 4, 5, 9, 10)
	getsecond(0, 4, 1, 2)

	res:=divide(2,1)
	fmt.Println(res)

}

func getsecond(a ...int) {
	sort.Ints(a)
	fmt.Println(a[len(a)-2])
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


what is the use of init function in golang