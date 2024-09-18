package main
import "fmt"
type person struct{
	name string
	age int
}

type animal struct{
	name string
}
type dog struct{
	animal
	breed string
}
func main(){

d:=dog{animal{"snoopy"},"desi"}
fmt.Println(d)

	p:=person{"harini",26}
	var p2 *person
	p2=new(person)

fmt.Println(p2)
	p1:=&person{"hari",10}
	fmt.Println(p1)
		
	fmt.Println(p)
    p.change1("friend")
	p1.change("fried")
	fmt.Println(p1)
	fmt.Println(p)
	(p1).change("budd")
	fmt.Println(p1)
	(&p).change("buddy")
	fmt.Println(p)
}

func (p *person)change(str string){
	(*p).name=str
}
func (x person)change1(str string) {
	x.name=str
}

//https://yourbasic.org/golang/structs-explained/