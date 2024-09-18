package main

import "fmt"
import "strings"

func main(){
	s  := "Hello HII life"

x:=strings.Split(s," ")
fmt.Println(x)
for _,b:=range x{
	r:=[]rune(b)
	
	for i,j:=0,len(r)-1;i<j;i,j=i+1,j-1{
       r[i],r[j]=r[j],r[i]
	}
	fmt.Print(string(r)+" ")
}
}