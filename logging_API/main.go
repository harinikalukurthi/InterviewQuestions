package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)
func main(){
router:=gin.Default()
gin.DebugPrintRouteFunc=func(httpMethod string, absolutePath string, handlerName string, nuHandlers int){
	log.Printf("endpoint %v %v %v %v",httpMethod,absolutePath,handlerName,nuHandlers)
}
//gin.DisableConsoleColor()
f,_:=os.Create("logging.log")
gin.DefaultWriter=io.MultiWriter(f,os.Stdout)

router.GET("/getdata",getdata)
router.Run()
}
func getdata(c *gin.Context){
	c.JSON(200,gin.H{
		"data":"I am from get Data",
	})
}