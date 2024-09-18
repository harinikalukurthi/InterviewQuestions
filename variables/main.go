package main

import "fmt"

//creating alias for packages

import fm "fmt"

var ay int = 10

const cs int = 6

func main() {
	var x int
	var _name string
	var f32 float32
	var f64 float64
	var b byte // synonym of uint8
	var r rune // which is equal to int32
	var u uint
	var y bool
	var p uintptr
	var c complex64
	_0 := 0
	fmt.Println(_0)
	__ := "hii"
	fmt.Println(__)
	fmt.Println("integer", x)
	fmt.Println("string", _name)
	fmt.Println("flaot32", f32)
	fmt.Println("float64", f64)
	fmt.Println("byte", b)
	fmt.Println("rune", r)
	fmt.Println("uint", u)
	fmt.Println("boolean", y)
	fmt.Println("unsigned ptr", p)
	fmt.Println("complex number", c)
	x = 10
	f32 = 7.8
	f64 = 90.76678
	_name = "harini"
	b = 10
	r = 100
	c = complex(6, 9)
	c1 := complex(2.7, 3.4)
	character := "very bad"
	fmt.Printf("integer %d and type %T\n", x, x)
	fmt.Printf("float %f and type %T\n", f32, f32)
	fmt.Printf("float %g and type %T\n", f32, f32)
	fmt.Printf("float %g and type %T\n", f64, f64) //the %g remove the 0 after the number
	fmt.Printf("float %f and type %T\n", f64, f64)
	fmt.Printf("string %s and type %T and length of the string %d\n", _name, _name, len(_name))
	fmt.Printf("byte %b and type %T\n", b, b)
	fmt.Printf("rune %d and type %T\n", r, r)
	fmt.Printf("complex %d and type %T\n", c, c)
	fmt.Printf("complex %d and type %T\n", c1, c1)
	fmt.Println("realnumber", real(c))
	fmt.Println("img number", imag(c))
	fmt.Println("combinigning two strings", _name+" "+character)

	//decleration of variables:
	var d int
	var a int = 20
	var t = 6.9
	j := 10
	d = 30
	k, l := 9, true
	var w, v int = 3, 4
	fmt.Println(d, j, a, t, w, v, k, l)

	a1, a2 := 9, 10
	a2, a3 := 34, 89
	fmt.Println(a1, a2)
	fmt.Println(a2, a3)
	//var 8ubv int  invalid
	/*For Types:
	  int, int8, int16, int32, int64, uint,
	  uint8, uint16, uint32, uint64, uintptr,
	  float32, float64, complex128, complex64,
	  bool, byte, rune, string, error
	*/

	//constants
	//constants cannot be declared using :=
	//const Pi:=3.14 invalid

	const w1 = 10
	const w2 int = 11
	// w1=16 cannot assign as the w1 is a constant..
	fmt.Println(w1, w2)
	var s1 = "i am very bad"
	var s2 = `i am so bad`
	fmt.Println(s1)
	fmt.Println(s2)
	var s3 = "Greetings and\n\"Salutations\""
	fmt.Println(s3)
	var s4 = `"Greetings and\n\"Salutations\""`
	fmt.Println(s4)

	var z int = 20
	f := float32(z)
	fmt.Println(z, f)
	const value = 10
	var i int = value
	var f1 float32 = value
	fmt.Println(i, f1)
	var b1 byte = 128
	fmt.Println(b1)

	// creating alias for data types
	type IZ int
	var a10 IZ = 1
	fm.Println(a10)

	

}
