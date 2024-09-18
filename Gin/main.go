package main

import (
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
)

type person struct {
	XMLName xml.Name `xml:"name"`
	FirstName string  `xml:"firstname"`
	LastName string   `xml:"lastname"`
}

func main() {
	router := gin.Default()
	router.GET("/Print", Print)
	router.GET("/Printname/:name", Printname)
	router.GET("/printPerson",printPerson)
	router.Run()

}

func Print(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello world"})
}

func Printname(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"message": "Hello " + name})
}

func printPerson(c *gin.Context){
c.XML(http.StatusOK,person{FirstName: "Harini",LastName: "Reddy"})
}
