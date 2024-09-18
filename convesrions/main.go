package main
import "fmt"
const y int=10 //explicit variable
const z=9.8 //implicit constant variable
const w=10/2 //constanst number must be caluculated before the compile time
//const g=getNumber() //this will through error bcz we have to get the constant number before compilation of the code
func main(){

//constants can only be declared for float,int,bool,strings
	fmt.Println(y,z)
	var x int=10
	var f float32=3.6
	fmt.Println("int value and float value",x,f)
	fmt.Println(int(f),float32(x))


	//enumerations

	const(
		i=iota
		a
		b
	)
	fmt.Println(i,a,b)
}
func getNumber()int{
	return 7
}