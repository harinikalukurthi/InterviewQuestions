package main

import (
	"americanExpress/automobiles"
	"fmt"
)

func main() {
	accord := automobiles.Car{
		Name: "Honda",
		Type: "sample",
		//horsePower: 1,
	}
	fmt.Println(accord)

}
