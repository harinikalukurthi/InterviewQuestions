package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router:=gin.New() #even with this also we start a router but this will not log our requests on the terminal
	//router.GET("/getData", getData)
	router.POST("/postData", postData)
	//router.GET("/queryData",queryData)
	router.GET("/getfields/:name/:age",getfields)
	//router.Run()
	//http.ListenAndServe(":9090",router)
	auth:=gin.BasicAuth(gin.Accounts{
		"harini":"123",
		"hari":"123",
	})

	admin:=router.Group("/admin",auth)
	{
		admin.GET("/getData", getData)
	}
	client:=router.Group("/clinet")
	{
		client.GET("/queryData",queryData)
	}
	server:= &http.Server{
		Addr: ":9090",
		Handler: router,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
	}
	server.ListenAndServe()
}
func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"GET Method data": "Hi This is GET method from GIN",
	})
}

func postData(c *gin.Context) {
	body := c.Request.Body
	value, err := io.ReadAll(body)
	if err != nil {
		fmt.Errorf("cannot read from request")
		return
	}
	c.JSON(200, gin.H{
		"POST Method data": "Hi This is POST method from GIN",
		"bosydata":         string(value),
	})
}

func queryData(c *gin.Context){
	name:=c.Query("name")
	age:=c.Query("age")
	c.JSON(200,gin.H{
		"data":"Hi I am querystring request call",
		"name": name,
		"age":age,
	})
}

func getfields(c *gin.Context){
	name:=c.Param("name")
	age:=c.Param("age")
	c.JSON(200,gin.H{
		"data":"Hi I am from get specific feilds method",
		"name": name,
		"age": age,
	})
}
