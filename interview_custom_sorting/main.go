package main

import (
	"fmt"
	"sort"
)

type Employee struct {
	Name    string
	Id      int
	Age     int
	Address string
}

func main() {
	e := []Employee{
		{"Harini", 1, 25, "Marala"},
		{"Harini", 2, 23, "Mudigubba"},
		{"Arthi", 12, 32, "Bengaluru"},
		{"Harish", 4, 13, "Chennai"},
	}
	fmt.Println(e)
	sort.Slice(e, func(i, j int) bool {
		if e[i].Name == e[j].Name {
			return e[i].Age < e[j].Age
		}
		return e[i].Name < e[j].Name
	})

	fmt.Println(e)

}
