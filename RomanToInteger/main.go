package main

import "fmt"

func main() {

	fmt.Println(romanToInt("XIV"))
	fmt.Println(romanToInt("III"))

}
func romanToInt(s string) int {
	m := make(map[rune]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	number := 0
	for i := 0; i < len(s); {
		v1, _ := m[rune(s[i])]
		if i < len(s)-1 {
			v2, _ := m[rune(s[i+1])]
			if v1 < v2 {
				number = number + (v2 - v1)

				i = i + 2
			} else {
				number = number + v1
				i = i + 1

			}
		}

	}
	return number

}
