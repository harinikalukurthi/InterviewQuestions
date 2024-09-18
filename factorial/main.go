package main
import "fmt"

func main(){
	n:=5
	fmt.Println(Fact(n))
}

func Fact(n int)int{
	if n<=1{
		return 1
	}
	return n*Fact(n-1)
}