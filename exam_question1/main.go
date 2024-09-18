package main
import "fmt"
import "strings"

func main(){
	str:="You can do it abcd the way you like"
	fmt.Println(Longeststringlength(str))
}
func Longeststringlength(str string)string{

	arr:=strings.Split(str," ")
   if len(arr)==1{
	return str
   }
	m:=make(map[string]int)
	for _,y:=range arr{
		m[y]=len(y)
	}
	fmt.Println(m)
  
   var x int 
   var res string
   for k,v :=range m{
		if v%2==0 && v>x{
			x=v
			res=k
		}
	}
	return res
}