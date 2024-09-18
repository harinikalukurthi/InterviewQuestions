package main
import "fmt"
func main(){
	s:="hello"
	sol:=count(s)
	fmt.Println(sol)

}
func count(s string)map[rune]int{
	x:=make(map[rune]int)
	for _,val:=range s{
		fmt.Println(val)
		if count,exist:=x[val];exist{
			x[val]=count+1
		}else{
			x[val]=1
		}
	}
	return x
}