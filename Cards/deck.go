package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func (d deck)print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck{
cards:=deck{}
cardSuits:=[]string{"hearts","diamonds","Spades","clubs"}
cardValues:=[]string{"ace","two","three","four"}
for _,suits:=range cardSuits{
	for _,values:=range cardValues{
		cards=append(cards,suits+"of" +values)
	}
}
return cards
}

func deal(d deck,handSize int) (deck,deck){
	return d[:handSize],d[handSize:]
}

func (d deck)savetoFile(filename string) error{
	string1:=strings.Join(d,",")
	return os.WriteFile(filename,[]byte(string1),0666)

}

func newDeckToFile(filename string)deck{
	 bs,err:=os.ReadFile(filename)
	 if err!=nil{
		fmt.Println(err)
		os.Exit(1)    
	}
    str:=string(bs)
	x:=strings.Split(str, ",")
	return deck(x)

}

func (d deck)shuffle(){
	for i:= range d{
          x:=rand.Intn(len(d)-1)
		  d[x],d[i]=d[i],d[x]
	}
}