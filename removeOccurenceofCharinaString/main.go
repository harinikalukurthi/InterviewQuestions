package main

import "fmt"

func main() {
	s := "Helllllo"
	k := 'l'
	fmt.Println(removeOcur(s, k))

}
func removeOcur(s string, k rune) string {

	res:=[]rune{}
	for _, v := range s {
		if v != k {
			res = append(res,v)
		}
	}
	return string(res)
}
