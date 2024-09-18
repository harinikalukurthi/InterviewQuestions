package main

import (
	"fmt"
)

/*
 * Complete the 'getHashedURL' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING url
 *  2. STRING hash_string
 *  3. INTEGER k
 */
var x = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', ':', '/', '.'}

func getHashedURL(url string, hash_string string, k int32) string {
	// Write your code here
	var result []byte

	s := split(url, k)
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		str := sum(s[i], len(hash_string))
		result = append(result, str)
	}

	return string(result)
}
func split(url string, k int32) []string {
	var pieces []string
	len1 := int(k)
	for i := 0; i < len(url); i += len1 {
		end := i + len1
		if end > len(url) {
			end = len(url)
		}
		pieces = append(pieces, url[i:end])
	}
	return pieces
}
func findIndex(char rune) int {
	for i, c := range x {
		if c == byte(char) {
			return i
		}
	}
	return -1
}
func sum(s string, y int) byte {
	sum := 0
	fmt.Println(s)
	for _, char := range s {
		index := findIndex(char)
		if index != -1 {
			sum += index
		}
	}
	fmt.Println("sum is %d", sum)
	val := sum % y
	fmt.Println("modulo %d", val)
	return x[val]
}

func main() {
	url := "https://xyz.com"
	hash_string := "hash_string"
	k := int32(4)
	hashedURL := getHashedURL(url, hash_string, k)
	fmt.Println(hashedURL)
}
