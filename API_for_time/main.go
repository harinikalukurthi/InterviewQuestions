package main

import (
	"net/http"
	"strings"
	"time"
)


func main(){
 http.HandlerFunc("/api/time", getTime)
 http.ListenAndServe("localhost:8080",nil)
}
func getTime(w http.ResponseWriter,r *http.Request){
response:=make(map[string]string)
tz:=r.URL.Query().Get("tz")
locations:=strings.Split(tz,",")

if len(locations)<1{
	
}

for _,loc:=range locations{

}
	
}



