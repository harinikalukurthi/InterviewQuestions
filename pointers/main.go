package main
import "fmt"
func main(){
	a:=10;
b:=&a
fmt.Println(a,b)
fmt.Println(a,*b)
a=a/2;
fmt.Println(a,b)
*b=*b/2
fmt.Println(a,b)
*b=34
fmt.Println(a,b)

}