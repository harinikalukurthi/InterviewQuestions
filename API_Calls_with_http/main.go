package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customers struct{
	Name string `json:"full_Name" xml:"name"`
	Id int      `json:"Id" xml:"id"`
	Zipcode string `json:"Zip_code" xml:"zipcode"`
}
func main(){
	mux:=http.NewServeMux()
	mux.HandleFunc("/greet",greet)
	mux.HandleFunc("/customerslist",getcustomers)
	log.Fatal(http.ListenAndServe("localhost:8000",mux))

}
func greet( w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"hello world")
}
func getcustomers(w http.ResponseWriter,r *http.Request){
	customer:=[]Customers{
		{"Harini",1,"12345"},
		{"Hari",2,"34521"},
	}


	if r.Header.Get("Content-Type")=="application/xml"{
	w.Header().Add("Content-Type", "application/xml")
	xml.NewEncoder(w).Encode(customer)
	}else{
		w.Header().Add("Content-Type", "application/json")
	   json.NewEncoder(w).Encode(customer)
	}
}