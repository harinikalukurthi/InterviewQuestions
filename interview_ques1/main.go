package main

import (
	"fmt"
	"sync"
)

const contact = 2800
const permanent = 3000

type employeetype string

const contracttype employeetype = "contract"
const permanenttype employeetype = "permanent"

type employee struct {
	employee string
	id       int
	days     int
	key      employeetype
}

var wg sync.WaitGroup

func main() {
	m := make(map[int]employee)

	m[2] = employee{"b", 2, 21, contracttype}
	m[3] = employee{"c", 3, 20, permanenttype}
	m[4] = employee{"b", 2, 21, contracttype}
	m[5] = employee{"c", 3, 20, permanenttype}
	m[6] = employee{"b", 2, 21, contracttype}
	m[7] = employee{"c", 3, 20, permanenttype}
	m[8] = employee{"b", 2, 21, contracttype}
	m[9] = employee{"c", 3, 20, permanenttype}
	m[10] = employee{"b", 2, 21, contracttype}
	m[11] = employee{"c", 3, 24, permanenttype}
	m[12] = employee{"b", 2, 21, contracttype}
	m[13] = employee{"c", 3, 20, permanenttype}
	ch := make(chan int, 10)
	wg.Add(10)
	for _, v := range m {
		go v.calcsalary(ch)
	}
	fmt.Println(<-ch)
	wg.Wait()
}

func (e employee) calcsalary(c chan int) {
	var salary int
	if e.key == "contract" {
		salary = (e.days * contact)
	} else {
		salary = (e.days * permanent)
	}
	wg.Done()
	c <- salary
}
