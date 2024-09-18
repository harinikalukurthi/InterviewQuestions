package main
import(
	"fmt"
)
func main(){
	fmt.Println("beuffered channels")
	ch:=make(chan int,3)
	ch<-1
	ch<-2
	ch<-3
	ch<-4
	num:=<-ch
	fmt.Println(num)
	num=<-ch
	fmt.Println(num)
	num=<-ch
	fmt.Println(num)

}