package main

import "fmt"

type person struct {
	fname   string
	lname   string
	contactInfo
}
type contactInfo struct {
	email   string
	zipcode int
}

func main() {
	p := person{fname: "Harini",
		lname: "reddy",
		contactInfo: contactInfo{
			email:   "harinireddy@gmail.com",
			zipcode: 515159,
		},
	}
	p.display()
	p.updateName()
	p.display()
	
}

func (p person) display(){
	fmt.Println(p.lname)
     fmt.Println(p.fname)
}
func (p *person)updateName(){
	p.fname="Hari"
}