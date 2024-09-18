package main
import "fmt"
type animal struct{
	name string
}
type cow struct{
	animal
	sound string
}
func main(){
c:=cow{animal{"baby"},"amba"}
fmt.Println(c)
}