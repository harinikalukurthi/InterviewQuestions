package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedlist struct {
	head *node
}

func (l *linkedlist) addlast(x int) {
	n := &node{x, nil}

	if l.head == nil {
		l.head = n
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = n
	return
}
func (l *linkedlist) addFirst(x int) {
	n := &node{x, nil}
	if l.head == nil {
		l.head = n
		return
	}
	n.next = l.head
	l.head = n
	return
}
func (l *linkedlist) Display() {
	if l.head == nil {
		fmt.Println("list is empty,nothing to display")
		return
	}
	current := l.head
	for current != nil {
		fmt.Print(current.data)
		fmt.Print("-->")
		current = current.next
	}
	fmt.Println()
}
func (l *linkedlist) DeleteFirst() {
	if l.head == nil {
		fmt.Println("list is empty,nothing to delete")
		return
	}
	current := l.head
	fmt.Println("deleted element", current.data)
	l.head = l.head.next
	current.next = nil
}
func (l *linkedlist) Deletelast() {
	if l.head == nil {
		fmt.Println("list is empty,nothing to delete")
		return
	}
	current := l.head
	last := l.head.next
	for last.next != nil {
		current = last
		last = last.next
	}
	current.next = nil
	fmt.Println("deleted element is", last.data)
}
func (l *linkedlist) Length() {
	count := 0
	current := l.head
	for current != nil {
		count = count + 1
		current = current.next
	}
	fmt.Println("length", count)

}

func (l *linkedlist)reverse(){
	if l.head==nil{
		fmt.Println("nothing to reverse")
		return
	}
	var prev *node=nil
	current:=l.head
    nextnode:=current.next

for nextnode!=nil{
current.next=prev
prev=current
current=nextnode
nextnode=nextnode.next
}
current.next=prev
l.head=current

}
func main() {
	l := linkedlist{}
	l.Display()
	l.addFirst(1)
	l.Display()
	l.addFirst(2)
	l.addFirst(3)
	l.Display()
	l.DeleteFirst()
	l.Display()
	l.Length()
	l.Deletelast()
	l.Display()
	l.addlast(4)
	l.addlast(5)
	l.Display()
	l.reverse()
	l.Display()
}
