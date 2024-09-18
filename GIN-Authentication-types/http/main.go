package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func main(){

	auth:=gin.BasicAuth(gin.Accounts{
		"harini":"123",
		"hari":"123",
	})

	router:=gin.Default()
	router.GET("/data",auth,getdata)
	router.Run()
}
func getdata(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"message":"Hello world"})
}

//can be called using
//curl -u harini:123 -X GET http://localhost:8080/data