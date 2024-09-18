package main


import "testing"

func TestFactorial(t *testing.T){
	val:=[]struct{
		input int
		output int
	}{
		{
			input: 5,
			output: 120,
		},{
			input: 0,
			output: 1,
		},{
			input:2,
			output:2,
		},
	}
	for _,v:=range val{
		if x:=Fact(v.input);x!=v.output{
			t.Errorf("expected %v and got %v",v.output,x)
		}
	}
}