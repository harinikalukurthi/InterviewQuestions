package main

import (
	"fmt"
	"sort"
)

func main(){
x:=[]int{7,8,3,1,9}
sorting(x)
fmt.Println(x[len(x)-2])
}
func sorting(x []int){
	sort.Slice(x,func(i,j int)bool{
		return x[i]<x[j]
	})
}