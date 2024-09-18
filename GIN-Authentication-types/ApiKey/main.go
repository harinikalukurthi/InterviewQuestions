package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var key int=123456789
func main() {
	router := gin.Default()

	router.GET("/data", getData)

	router.Run()
}
func getData(c *gin.Context) {

	apikey:=c.Request.Header.Get("APIKEY")
	key2,_:=strconv.Atoi(apikey)
	if key==key2{
		c.JSON(http.StatusOK,gin.H{"message":"Hello world"})
		return
	}
	c.JSON(http.StatusUnauthorized,gin.H{"message":"unauthorized"})

}
