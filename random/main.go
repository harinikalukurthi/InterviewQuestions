package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

//var letters = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
var letters=[]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
func main() {
	s := []string{}
	ru:=[]rune{}
	for i:=0;i<2;i++{
	r := rand.Intn(len(letters))
	s = append(s, string(letters[r]))
	ru=append(ru, rune(letters[r]))
	}
	for i:=0;i<3;i++{
	n := rand.Intn(10)
	x:=strconv.Itoa(n)
	s = append(s, x)
	ru=append(ru,int32('0'+n))
	}
	fmt.Println(s)
	fmt.Println(string(ru))
}