package main
import "fmt"
type nums interface{
       sum() int
}

type sample struct{
	x int
    y int
}
func (s *sample)sum()int{
	if s==nil{
		fmt.Println("empty")
	}
	fmt.Println(s.x+s.y)
	return s.x + s.y
}

type number int32

func (i number) sum() int32 {
	if i>0{
	return int32(i)
	}
	return int32(-i)
}

func main(){
	f:=&sample{}
	fmt.Printf("%v , %T\n",f,f)
	fmt.Println("interfaces")
	s:=&sample{3,4}
	fmt.Println(s.sum())
	x:=number(1)
	fmt.Println(x)
	fmt.Printf("%v , %T\n",x,x)
	var n1 nums=&sample{5,4}
	fmt.Println(n1.sum())
    fmt.Printf("%v , %T\n",s,s)

var n nums
fmt.Printf("%v , %T\n",n,n)
//n.sum()  //Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
var st *sample
n=st
//n.sum()
fmt.Printf("%v , %T\n",n,n)

var i interface{}
describe(i)
i="hello"
describe(i)
i=9.0
describe(i)

var k interface{} = "hello"

	s1 := k.(string)
	fmt.Println(s1)

	s1, ok := k.(string)
	fmt.Println(s1, ok)

	f1, ok := k.(float64)
	fmt.Println(f1, ok)

	f1,_ = k.(float64) // panic
	fmt.Println(f1)
}
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
