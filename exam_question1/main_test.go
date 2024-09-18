package main

import "testing"

func TestLongeststringlength(t *testing.T) {
	s := []struct {
		input  string
		output string
	}{
		{
			input:  "I am good girl",
			output: "good",
		},
		{
			input:  "K",
			output: "K",
		},
	}

	for _,v:=range s{
		if x:=Longeststringlength(v.input);x!=v.output{
			t.Errorf("expected %v but got %v",v.output,x)
		}
	}
}
