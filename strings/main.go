package main

import (
	"fmt"
	"strings"
)
func main(){
	fmt.Println("hello world")
	fmt.Println("I am \n bad")
	fmt.Println(`i am string`)
	fmt.Println(`i am \n doll`)

	str:="marala"
	b:=str[2]
	fmt.Println("value of b",rune(b))
	c:=str[:4]
	fmt.Println("value of c",c)
	fmt.Println("length of string",len(str))
	//str[1]='w' //strings are immutable we cant change them
	fmt.Println(str)
	for i,v:=range str{
		fmt.Printf("index %d and value %c\n",i,v)
	}
for _,v:=range str{
	fmt.Printf("index %c and value %v\n",v,v)
}

myslice1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}
  fmt.Println(myslice1)
    // Creating a string from the slice
    mystring1 := string(myslice1)
  
    // Displaying the string
    fmt.Println("String 1: ", mystring1)


	string1:="@welcome my girl!$#"
	res:=strings.Trim(string1,"@#$^!")
	fmt.Println(res)
	fmt.Println(strings.TrimLeft(string1,"@"))
	fmt.Println(strings.TrimRight(string1,"#$"))

	string2:="Hello all, come to my world, where you will be tortured,killed,laughed,loved"
	fmt.Println(strings.Split(string2,","))
	fmt.Println(strings.Compare("gfg", "Geeks"))
     
    fmt.Println(strings.Compare("GeeksforGeeks", 
                               "GeeksforGeeks"))
     
    fmt.Println(strings.Compare("Geeks", " GFG"))
     
    fmt.Println(strings.Compare("GeeKS", "GeeKs"))
	x:=strings.Contains("BLOCKCOMMITr018-2209149e-d1c0-4373-aa8e-d05bb9da2f2c", "BLOCKCOMMIT")
	fmt.Println("contains",x)

}
