package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var mu = sync.Mutex{}

type data struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Due_date    string `json:"due_data"`
}

var counter = 0
var m = make(map[int]data)
var d=make(map[string]string)

func main() {
	d["harini"]="1234"
	router := gin.Default()
	router.POST("/adddata", add)
	router.GET("/list", list)
	router.PUT("/update/:id",updatetask)
	router.DELETE("/delete/:id",deletetask)
	router.GET("/get/:id",getid)
	router.PUT("/:username/pass",changePass)


	router.Run("localhost:8080")

}
func changePass(c *gin.Context){
	user_name:=c.Param("username")
	pass:=c.Param("pass")

	if val,ok:=d[user_name];ok{
		if val!=pass{
			d[user_name]=pass
			c.JSON(http.StatusOK,gin.H{"message":"password is updated"})
			return
		}else{
			c.JSON(http.StatusBadRequest,gin.H{"message":"please eneter another password"})
			return
		}
	}
	c.JSON(http.StatusNotFound,gin.H{"message":"user is not found"})
}
func add(c *gin.Context) {
	var request data
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mu.Lock()
	defer mu.Unlock()
	counter++
	m[counter] = request
	c.JSON(http.StatusOK, m[counter])
}
func list(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	c.JSON(http.StatusOK, m)
}
func deletetask(c *gin.Context) {
	idstr := c.Param("id")
	id,err := strconv.Atoi(idstr)
	if err!=nil{
		c.JSON(http.StatusBadRequest,err)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	delete(m, id)
	c.JSON(http.StatusOK, m)
}

func updatetask(c *gin.Context){
	var request data
	idstr:=c.Param("id")
	id,err:=strconv.Atoi(idstr)
	if err!=nil{
		c.JSON(http.StatusBadRequest,err)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	if err:=c.ShouldBindJSON(&request);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingData:=m[id]
	existingData.Title=request.Title
	existingData.Description=request.Description
	existingData.Due_date=request.Due_date
	m[id]=existingData

	c.JSON(http.StatusOK,m[id])
}
func getid(c *gin.Context){
	idstr:=c.Param("id")
	id,err:=strconv.Atoi(idstr)
	if err!=nil{
		c.JSON(http.StatusBadRequest,err)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	c.JSON(http.StatusOK,m[id])
}
