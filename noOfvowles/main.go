package main
import "fmt"
func main(){
s:="hello"
countVol:=0
countcon:=0
for _,x:=range s{
	if x=='a'||x=='e'||x=='i'||x=='o'||x=='u'{
       countVol++
	}else{
		countcon++
	}
}
fmt.Println(countVol,countcon)
}