package main
import(
	"fmt"
)
type shape interface{
	getArea() float64
}
type traingle struct{
	height float64
	base float64

}
type square struct{
	sidelength float64
}
func(t traingle)getArea()float64{
return t.height*t.base
}
func(s square)getArea() float64{
	return s.sidelength*s.sidelength
}

func PrintArea(s shape){
fmt.Println(s.getArea())
}
func main(){
	s:=square{sidelength: 10}
	t:=traingle{height: 10,base: 10}
	PrintArea(s)
	PrintArea(t)

}