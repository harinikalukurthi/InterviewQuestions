package main
import "fmt"
func main(){

	s:="abcbaf"

	fmt.Println(palin(s))
}
func palin(s string) string{
	r:=[]rune(s)
	for i,j:=0,len(r)-1;i<j;i,j=i+1,j-1{
        if r[i]!=r[j]{
			return "not a palindrome"
		}
	}
	return "palindrome"
}