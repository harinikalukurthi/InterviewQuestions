package main
import "fmt"
import "strings"
import(
	"bytes"
 "reflect"
 "sort"
)
func main(){
	var a []int
	a=append(a,1)
	a=append(a,2,3)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	b:=[]int{1,2,3}
	fmt.Println(b,len(b),cap(b))
	b=append(b,9,8,7)
	fmt.Println(b,len(b),cap(b))
	var arr =[]int{10,11,12}
	b=append(b,arr...)
	fmt.Println(b,len(b),cap(b))
	b=b[1:]
	fmt.Println(b,len(b),cap(b))
	c:=b[:len(b)-1]
	fmt.Println(c)
	c=b[3:5]
	fmt.Println(c)

	//changing the vaue of an slice will change its underlying value in array
s:=[9]int{1,2,3,4,5,6,7,8,9}
fmt.Println(s)
f:=s[2:7]
fmt.Println(f)
f[3]=11
fmt.Println(f)
fmt.Println(s)

stru:=[]struct{
i int
b bool
}{
	{1, true},
	{2,false},
	{3,true},
	{4,false},
}
fmt.Println(stru)

//zero value of a slice
var sli []int
fmt.Println(sli)
if sli==nil{
	fmt.Println("nil")
}
//creating the slice  with make we can specify only length or we can specify both len and cap

s1:=make([]int,2,2)
fmt.Println(s1)

s2:=make([]int,3)
fmt.Println(s2)

//creating 2 dimenstional slice
var s3=[][]string{
	[]string{"a","v","y"},
	[]string{"g","p","i"},
	[]string{"e","u","r"},
}
fmt.Println(s3)
fmt.Println(len(s3))
for i := 0; i < len(s3); i++ {
	fmt.Printf("%s\n", strings.Join(s3[i], " "))
}


for d,t:=range a{
	fmt.Println(d,t)
}
//comparing both the slices
slice1:=[]int{2,3,4}
slice2:=[]int{2,3,4}
if slice1==nil{
	fmt.Println("slice1 is nil")
}else{
	fmt.Println("slice1 not nil")
}
if slice2==nil{
	fmt.Println("slice2 is nil")
}else{
	fmt.Println("slice2 is not nil")
}

fmt.Println(reflect.DeepEqual(slice1,slice2))

fmt.Println(equal(slice1,slice2))
//if slice1==slice2{
	//fmt.Println("equal")
//}
slice_1 := []byte{'!', '!', 'G', 'e', 'e', 'k', 's', 'f',  
                 'o', 'r', 'G', 'e', 'e', 'k', 's', '#', '#'}
     
    slice_2 := []byte{'*', '*', 'A', 'p', 'p', 'l', 'e', '^', '^'}
     
    slice_3 := []byte{'%', 'g', 'e', 'e', 'k', 's', '%'}
 
    // Displaying slices
    fmt.Println("Original Slice:")
    fmt.Printf("Slice 1: %s", slice_1)
    fmt.Printf("\nSlice 2: %s", slice_2)
    fmt.Printf("\nSlice 3: %s", slice_3)
 
    // Trimming specified leading 
    // and trailing Unicodes points
    // from the given slice of bytes
    // Using Trim function
    //res1 := bytes.Trim(slice_1, "!#")
    res2 := bytes.Trim(slice_2, "*^")
    res3 := bytes.Trim(slice_3, "@")
	res1 := bytes.Trim([]byte("****Welcome to GeeksforGeeks****"), "*")

    // Display the results
    fmt.Printf("New Slice:\n")
    fmt.Printf("\nSlice 1: %s", res1)
    fmt.Printf("\nSlice 2: %s", res2)
    fmt.Printf("\nSlice 3: %s\n", res3)
	byteSlice := []byte("  Hello, World!  ")
    trimmed := bytes.TrimSpace(byteSlice)
    fmt.Println(string(trimmed))

	//sorting slice using inbuilt method
	intSlice := []int{5, 2, 6, 3, 1, 4}
    sort.Ints(intSlice)
    fmt.Println(intSlice)

	//sorting the slice without inbuilt method
	a1 := []int{5, 3, 4, 7, 8, 9}
sort.Slice(a1, func(i, j int) bool {
    return a1[i] < a1[j]
})
for _, v := range a1 {
    fmt.Println(v)
}
}
func equal(x,y []int)bool{
	for i,_:=range x{
		if x[i]!=y[i]{
			return false
		}
	}
	return true
}