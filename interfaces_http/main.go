package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("erros", err)
		os.Exit(1)
	}*/

	//fmt.Println(resp)
	//os.Stdout(resp)
	//bs:=make([]byte,99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))
	fmt.Println(os.Args[1])
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error,err")
	}

	io.Copy(os.Stdout,file)
}
