package main

import "sort"
import "fmt"

func main() {
	a := []int{2, 5, 1, 7, 8}
	k := 4
	fmt.Println(search(a, k))
}
func search(x []int, key int) int {
	y := sorting(x)
	low := 0
	high := len(y) - 1
	mid := (low + high) / 2

	for low < high {

		if key > y[mid] {
			low = mid + 1
		} else if key < y[mid] {
			high = mid - 1
		} else {
			return mid
		}
		mid=(low+mid)/2
	}
	if low>high{
		return -1
	}
	return -1
}


func sorting(x []int) []int {
	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})
	return x
}
