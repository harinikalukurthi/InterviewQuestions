package main
import "fmt"
func main(){
	 n:=10
	fmt.Println(prime(n))
}
func prime(n int)bool{
	if n==0||n==1{
		return false
	}else if n==2{
		return true
	}
	for i:=2;i<n;i++{
		if n%i==0{
			return false
		}
	}
	return true
}