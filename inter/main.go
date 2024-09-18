package main

import "fmt"

type bot interface {
	getReading() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printmsg(eb)
	printmsg(sb)

}
func (eb englishBot) getReading() string {
	return "Hi There"
}
func (sb spanishBot) getReading() string {
	return "Hola"
}
func  printmsg(b bot) {
	fmt.Println(b.getReading())
}
