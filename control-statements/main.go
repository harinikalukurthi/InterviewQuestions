package main

import (
	"fmt"
	"time"
)

func main() {

	//if and else statement
	var x = 10
	if x == 10 {
		fmt.Printf("value of x is %d\n", x)
	} else {
		fmt.Println("value is not equal")
	}

	//nested if sample
	a, b := 13, 90
	if a == 13 {
		if b == 90 {
			fmt.Println("both values are equal")
		} else {
			fmt.Println("b is not equal")
		}
	} else {
		fmt.Println("a is not equal")
	}

	//if-else-if sample

	c := 100
	if c < 100 {
		fmt.Println("c is less than 100")
	} else if c > 100 {
		fmt.Println("c is greater than 100")
	} else {
		fmt.Println("c is equal to 100")
	}

	//for loop
	for i := 0; i < 3; i++ {
		fmt.Printf("the value of i is %d\n", i)
	}
	//while loop
	y := 0
	for y < 3 {
		fmt.Printf("the value of y is %d\n", y)
		y++
	}
	//infinite loop
	/*for{
		fmt.Println("infinite loop")
	}*/

	//using range
	q := []int{1, 2, 3, 4, 5}
	for r, t := range q {
		fmt.Println(r, t)
	}
	//iterating over a string

	var str = "i am very bad"
	for o, p := range str {
		fmt.Println(o, p)
	}

	//range over map

	map1 := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	for k, l := range map1 {
		fmt.Println(k, l)
	}

	// range over a channel
	ch := make(chan int)
	go func() {
		ch <- 10
		ch <- 100
		ch <- 1000
		ch <- 10000
		close(ch)
	}()
	for g := range ch {
		fmt.Println(g)
	}

	//switch samples

	switch day := 4; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("thursday")
	case 5:
		fmt.Println("friday")
	case 6:
		fmt.Println("saturday")
	case 7:
		fmt.Println("sunday")
	default:
		fmt.Println("invalid")
	}

	var xx = 1
	switch {
	case xx == 1:
		fmt.Println("one")
	case xx == 2:
		fmt.Println("two")
	case xx == 3:
		fmt.Println("three")
	}

	var l rune = 'a'

	switch l {
	case 'a':
		fmt.Println("its is A")
	case 'b':
		fmt.Println("its is B")
	case 'c':
		fmt.Println("its is C")
	case 'd':
		fmt.Println("its is D")
	}

	var st string = "good"
	switch st {
	case "good", "best":
		fmt.Println("super")
	case "bad", "worst":
		fmt.Println("ugly")
	}

	var value interface{}
	switch q := value.(type) {
	case bool:
		fmt.Println("value is of boolean type")
	case float64:
		fmt.Println("value is of float64 type")
	case int:
		fmt.Println("value is of int type")
	default:
		fmt.Printf("value is of type: %T", q)

	}
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Tuesday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
