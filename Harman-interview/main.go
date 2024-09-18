/*var x int
x=10

var x int, y string

x,y:=1,"abc"

var y *int
 z := new(int)
 what is the difference when we create a pointer with * and with new operator
 y is declared as a pointer to an int, but it is not initialized. By default, it is set to nil.
z is a pointer to an int that is initialized by the new(int) function, which allocates memory for an int value and returns a pointer to that memory. The memory is zeroed, so the value of the int is 0.


 polymorphism -using 
 func add(x ...int){

 }

 type person struct{
	name string
	Age

 }
 type Age struct{

 }
 s:=make([]int,18,18) (19,36) header,
 ch:=make(chan string)
 ch:=make(chan int,10)
 m:=make(map[int]string)

 func(){
	defer func(){
		if r:=recover();r!=nil{

		}

	}()
	panic()
 }


 a
 b and c

 inside
 import "a/b"

 type linkedlist struct{
	head *node
 }

 type node struct{
     data int
	 next *node
 }
 type tree struct{
	 head *node
 }
 type node struct{
	data int
	lNode *node
	Rnode *node
 }
 b-
 B
 func(n node)add(){

 }

func add(){

}

lets say we have struct with 3 field
PUT//what will happen if send a fourth field data into the  PUT, what will happen.
ans: The Go json package will ignore any fields that are present in the JSON but do not exist in the struct. So, the age field will be ignored without causing any error.

what are the oops concepts that we can implement with GO
1. polymorphism with interfaces
2. Inheritence with structures(type embedding)
3. abstraction through interfaces
4. Encapsulation with first letter captial and smaller

func testAdd(t *testing.T){
	test:=struct{
		inputx  int
		inputy  int
		expectresult iny
	}{
		inputx=2,
		inputy=3,
		expectresult=5,
	}
}
	if res:=sum(test.inputx+test.inputy);res!=test.expctedresult{
	t.Errorf("expected %v but got %v",expctedresult,res)
	}}

git checkout -b <new-branch>
git push <origin> <current-branch>

CMD[""]

difference betwen go.mod and go.sum
what is goroot and gopath
what are rules for packages(package imports)
authentication:process of verifying the identity of a user, device, or system. It determines who you are. looging to amazon prime
authorization:process of determining whether a user, device, or system has permission to access a resource or perform an action. checking whether we have prime subscription
deadlocks types
what will happen if we insert more data into channel size

kubectl get pods -n <name-space>
kubectl config set-context <context-name>*/

/* two variables string a="Silent", b:="Listen" */

package main

import (
	"fmt"
	"strings"
	"sort"
	"reflect"	
)

func main() {
	a := "Silent"
	b := "Listen"
	c := strings.ToLower(a)
	d := strings.ToLower(b)
	fmt.Println(check(c, d))
}
func check(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	x:=sort1(a)
	y:=sort1(b)
	return reflect.DeepEqual(x,y)
}

func sort1(a string)([]string) {
	s:=[]string{}
	//s:=[]rune(a)
	/*for i:=0;i<len(s);i++{
		for j:=0;j<len(s);j++{
			if s[i]>s[j]{
				s[i],s[j]=s[j],s[i]
			}
		}
	}*/
	for _,v:=range a{
		s=append(s,string(v))
	}
	sort.Strings(s)
	return s
}
