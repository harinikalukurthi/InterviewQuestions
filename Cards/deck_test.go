package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16,but got %v", len(d))
	}
	if d[0]!= "heartsoface"{
		t.Errorf("Expected first value is heartsoface,but got %v", d[0])
	}
	if d[len(d)-1]!= "clubsoffour"{
		t.Errorf("Expected last value is clubsoffour,but got %v", d[len(d)-1])

	}
}
func TestsaveToFileAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")
	d:=newDeck()
	d.savetoFile("_decktesting")
	loadDeck:=newDeckToFile("_decktesting")
if len(loadDeck)!=16{
	t.Errorf("Expected l16 cards,but got %v", len(d))
}
os.Remove("_decktesting")

}
