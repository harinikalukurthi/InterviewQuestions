package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Fruits struct {
	Name string
	Cost int
	City string
}

var fruit = []Fruits{}

func main() {
	router := gin.Default()
	router.POST("/newfruit", addfruit)
	router.GET("/list", listfruits)
	router.DELETE("/:name",deletefruitbyname)
	router.GET("/:name",listparticularfruit)
	router.PUT("/:name",updatecost)

	router.Run("localhost:8080")

}

func addfruit(c *gin.Context) {
	var newdata Fruits
	if err := c.ShouldBindJSON(&newdata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fruit = append(fruit, newdata)
	c.JSON(http.StatusCreated, newdata)
}

func listfruits(c *gin.Context) {
	c.JSON(http.StatusOK, fruit)
}

func deletefruitbyname(c *gin.Context){
    f:=c.Param("name")
	for i,x:=range fruit{
     if f==x.Name{
       fruit=append(fruit[:i],fruit[i+1:]...)
	   c.JSON(http.StatusOK, f )
	   return
	 }
	}
	c.JSON(http.StatusNotFound,gin.H{"message":"fruit not found to delete"})
}

func listparticularfruit(c *gin.Context){
	f:=c.Param("name")
	for _,x:=range fruit{
		if x.Name==f{
         c.JSON(http.StatusFound,x)
		 return
		}
	}
	c.JSON(http.StatusNotFound,gin.H{"message":"fruit not found to list"})
}

func updatecost(c *gin.Context){
	f:=c.Param("name")
	var cost struct {
		Cost int `json:"cost"`
	}
	if err := c.ShouldBindJSON(&cost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
for i,x:=range fruit{
	if x.Name==f{
		fruit[i].Cost=cost.Cost
		c.JSON(http.StatusOK,gin.H{"message":"fruit is updated "})
		return
	}
}
c.JSON(http.StatusNotFound,gin.H{"message":"Fruit not found to update cost"})

}