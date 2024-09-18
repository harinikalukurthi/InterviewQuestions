package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, -1, 2, -3, 1}
	target := 4
	fmt.Println(find_3(arr,target))
	fmt.Println(check2_efficient(arr, target))

	fmt.Println(find(arr, target))
	fmt.Println(check(arr, target))
	fmt.Println(find(arr, target))

}

func check2_efficient(x []int, t int)(int,int){ //using hash table
	m:=make(map[int]int)
	for i:=0;i<len(x);i++{
		sub:=t-x[i]
		if y,v:=m[sub];v{
			return y,i
		}else{
			m[x[i]]=i
		}
	}
	return 0,0
}

func find(x []int, t int) (int, int) {
	for i := 0; i < len(x); i++ {
		for j := i + 1; j < len(x); j++ {
			if x[i]+x[j] == t {
				return i, j
			}
		}
	}
	return 0, 0
}

func check(x []int, t int) (int, int) {
	sort.Slice(x,func(i,j int)bool{
		return x[i]<x[j]
	})
	for i,j:=0,len(x)-1;i<j;{
		if x[i]+x[j]==t{
			return i,j
		} else if x[i]+x[j]<t{
			i=i+1
		}else{
			j=j-1
		}

	}
	return 0,0
}

func find_3(x []int, t int) (int, int,int) {
	for i := 0; i < len(x); i++ {
		for j := i + 1; j < len(x); j++ {
			for k := j + 1; k < len(x); k++ {
			if x[i]+x[j]+x[k] == t {
				return i, j,k
			}
		}
		}
	}
	return 0, 0,0
}

