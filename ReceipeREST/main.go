package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json: "name"`
	Tags         []string  `json: "tags"`
	Ingredients  []string  `json: "ingredients"`
	Instructions []string  `json: "instructions"`
	PublishedAt  time.Time `json: "publishAt"`
}

var recpies []Recipe

func main() {
	router := gin.Default()
	router.POST("/recpies", NewRecipeHandler)
	router.GET("/recpies", GetRecpies)
	router.PUT("/recpies/:id", UpdateRecpie)
	router.DELETE("/recpies/:id", DeleteRecpie)
	router.GET("/recpiesbytags",Searchtags)
	router.Run()
}
func NewRecipeHandler(c *gin.Context) {
	var request Recipe
	    // Unmarshalling: converts JSON received in request body to Go struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	request.PublishedAt = time.Now()
	request.ID = xid.New().String()
	recpies = append(recpies, request)
	c.JSON(http.StatusOK, request)

}
func GetRecpies(c *gin.Context) {
		    // marshalling: converts object received to JSON

	c.JSON(http.StatusOK, recpies)

}
func UpdateRecpie(c *gin.Context) {
	id := c.Param("id")
	var recpie Recipe
	if err := c.ShouldBindJSON(&recpie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	index := -1
	for i := 0; i < len(recpies); i++ {
		if recpies[i].ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not find the id"})
	}
	recpies[index].ID = id
	recpies[index].Name = recpie.Name
	recpies[index].Ingredients=recpie.Ingredients
	recpies[index].Instructions=recpie.Instructions
	recpies[index].PublishedAt=time.Now()
	


	c.JSON(http.StatusOK, gin.H{"successfully updated": recpies[index]})

}
func DeleteRecpie(c *gin.Context) {
	id := c.Param("id")

	index := -1
	for i := 0; i < len(recpies); i++ {
		if recpies[i].ID == id {
			index = i
			break
		}
	}
	if index==-1{
		c.JSON(http.StatusBadRequest,gin.H{"could not find the id":id})
		return
	}
	recpies = append(recpies[:index],recpies[index+1:]...)
	c.JSON(http.StatusOK,gin.H{"successfully deleted":id})

}
func Searchtags(c *gin.Context){
	tag := c.Query("tag")
	listOfRecipes := make([]Recipe, 0)
	for _,recipe:=range recpies{
		for _,t:=range recipe.Tags{
			if t==tag{
				listOfRecipes = append(listOfRecipes, recipe)
			}
		}
	}
	if len(listOfRecipes)==0{
		c.JSON(http.StatusNotFound,gin.H{"tags were not found":tag})
		return
	}
	c.JSON(http.StatusOK,listOfRecipes)


}