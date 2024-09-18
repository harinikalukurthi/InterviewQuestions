package main

import "testing"

func TestAnagram(t *testing.T){
	anams:=[]struct{
		inputstr1 string
        inputstr2 string
		expectoutput bool
	}{
        {
			inputstr1: "listen",
			inputstr2: "silent",
			expectoutput: true,
		},
		{
			inputstr1: "abc",
			inputstr2: "bac",
			expectoutput: true,
		},
	}

	for _,test:=range anams{
		if x:=Anagram(test.inputstr1,test.inputstr2);x!=test.expectoutput{
			t.Errorf("expected %v and got %v",test.expectoutput,x)
		}
	}
}