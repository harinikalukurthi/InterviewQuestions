package main
import "fmt"
func main(){
	var a [4]int
	fmt.Println(a)
	a[0]=1
	a[1]=2
	a[2]=3
	a[3]=4
	fmt.Println(a)
	fmt.Println(len(a))
	var b=[3]int {7,8,9}
	fmt.Println(b)
	c:=[2]int{5,6}
	fmt.Println(c)
	a[3]=10
	fmt.Println(a)


	//comparing the arrays
	x:=[2]int{1,2}
	y:=[2]int{1,2}
	if x==y{
		fmt.Println("equal")
	}

}