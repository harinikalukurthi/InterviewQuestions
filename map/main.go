package main

import "fmt"

func main() {
	c := make(map[string]int)
	c["one"] = 1
	c["two"] = 2
	fmt.Println(c)
	m:=map[int][]string{
		1: []string{"one","One"},
	}
	fmt.Println(m)

	c["one"]=3
	fmt.Println(c)

	for k,v:=range c{
		fmt.Println(k,v)
	}

	fmt.Println( c["one"])
	delete(c,"four")
	fmt.Println(c)
}
