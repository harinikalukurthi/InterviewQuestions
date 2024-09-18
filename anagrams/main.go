package main

import "sort"
import "fmt"

func main(){
	str1:="tar"
	str2:="rac"
	fmt.Println(Anagram(str1,str2))
}
func Anagram(x,y string)bool{
	if len(x)!=len(y){
		return false
	}
	a:=sorting(x)
	b:=sorting(y)
	for i:=0;i<len(a);i++{
		if a[i]!=b[i]{
			return false
		}
	}
	return true
}

func sorting(s string)string{
r:=[]rune(s)
sort.Slice(r,func(i,j int)bool{
	return r[i]<r[j]
})
return string(r)
}