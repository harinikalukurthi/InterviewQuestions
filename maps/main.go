package main
import "fmt"
type areas struct{
   x , y string
}

func main(){
	fmt.Println("maps")

var m =make(map[string]areas)
m["bangalore"]=areas{
	"varthur", "domlur",
}

var n = map[string]areas{
"bangaluru":areas{
	"varthur","domlur",
},
"hyderabad":areas{
	"kondapur","ameerpet",
},
}

var o = map[string]areas{
"bangaluru":{"varthur","domlur"},
"hyderabad":{"ameerpet","kondapur"},
}

var a = make(map[int]string)
//inserting a key and value into map
a[1]="one"
fmt.Println(a)
a[1]="onee"
//retrieving a specific value
fmt.Println(a[1])
//deleting a value in map
delete(a,1)
fmt.Println(a)
a[2]="two"
val,ok:=a[2]
if ok{
	fmt.Println(val)
}

fmt.Println(m)
fmt.Println(n)
fmt.Println(o)


map1:=make(map[int]string)
map1[1]="one"
map1[2]="two"
map1[3]="three"
map1[4]="four"

map2:=make(map[int]string)
for k,v:=range map1{
	map2[k]=v
}
fmt.Println(map2)

}