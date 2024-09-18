package main

import "fmt"

func main() {
	s := "ABC"
	str := reverse(s)
	fmt.Println(str)
	count := len(str) - 1
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]), count)
		count--
	}

}
func reverse(s string) string {
	r := []rune{}
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, rune(s[i]))
	}

	return string(r)
}