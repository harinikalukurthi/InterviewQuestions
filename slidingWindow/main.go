package main

import "fmt"

func main() {
	a := []int{2, 3, 1, 4, 9, 7}
	fmt.Println(maxsum(a, 2))
	fmt.Println(desiresum(a, 5))

}
func maxsum(a []int, k int) int {
	sum := 0
	cursum := 0
	for i := 0; i < k; i++ {
		cursum += a[i]
	}
	sum = cursum

	for i := k; i < len(a); i++ {
		cursum = cursum + a[i] - a[i-k]
		if cursum > sum {
			sum = cursum
		}
	}
	return sum
}

func desiresum(a []int, desire int) [][]int {
	var result [][]int
	sum := 0
	start := 0

	for end := 0; end < len(a); end++ {
		sum = sum + a[end]
		if sum > desire {
			sum = sum - a[start]
			start++
		}
		if sum == desire {
			result = append(result, a[start:end+1])
		}
	}
	return result
}

func nonrepeatString(a string) []string {
	start := 0
	var s []string
	var end int
	m := make(map[string]int)
	for end = 0; end < len(a); end++ {
		_, ok := m[string(a[end])]
		if ok {
			start = end
		} else {
			m[string(a[end])] = 1
		}
	}
	return (s[start : end+1])
}
