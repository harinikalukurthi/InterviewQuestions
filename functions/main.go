package main
import "fmt"
func main(){
fmt.Println(sum(3,4))
p,q:=9,2
k,l:=swap(p,q)
fmt.Println(k,l)
fmt.Println(p,q)
w,u:=swapbyref(&p,&q)
fmt.Println(w,u)
fmt.Println(p,q)
fmt.Println(sumvar())
fmt.Println(sumvar(2))
fmt.Println(sumvar(8,6,5))
fmt.Println(sumvar(3,0,8,2,1,3))
//passing the slice as an input to the vardiac function
s:=[]int{1,2,3,4,5}
fmt.Println(sumvar(s...))
//anonymous function
func(){
	fmt.Println("hello anonymous")
}()
//assigning the anonymous function to a variable
val:=func(){
	fmt.Println(3+4)
}
val()

//passing the variable to a anonymous function
func(x,y int){
	fmt.Println(x+y)
}(9,7)

//passing anonymous function as a input to another function
value:=func(p,q string)string{
	return p+q+"Go"
}
GFG(value)
val1:=sam()
fmt.Println(val1("hi","namasthe"))
}
 
func GFG(i func(p,q string)string){
	fmt.Println(i ("Geeks" , "good"))
}

func sam() func(i,j string)string{
	myfun:=func(i,j string)string{
		return i+j+"hello"
	}
	return myfun
}

func sum(a int,b int)(int){
return a+b
}

//call by value
func swap(x int,y int)(a,b int){
	a,b=y,x
	return 
}

//call by reference
func swapbyref(x,y *int)(int,int){
*x,*y=*y,*x
return *x,*y
}

//variadic function: a function can accept any no.of arguements
func sumvar(x... int)(int){
	sum:=0
	for _,a:=range x{
sum=sum+a;
	}
	return sum
}

