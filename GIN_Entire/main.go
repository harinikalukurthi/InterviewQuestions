package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	//router.Use(Authenticator)
	admin:=router.Group("/admin",Authenticator)
	{
		admin.GET("/getData2", getData2)
		admin.GET("/getData3", getData3)
	}
	router.GET("/getData",getData)
	router.Run()

}
func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi I am getData",
	})
}
func getData2(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi I am getData2",
	})
}
func getData3(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi I am getData3",
	})
}

func Authenticator(c *gin.Context){
	if !(c.Request.Header.Get("Token")=="auth"){
		c.AbortWithStatusJSON(500,gin.H{
			"Message":"Token not present",
		})
		return
	}
	c.Next()
}
/*func Authenticator() gin.HandlerFunc {
	return func(c *gin.Context){
		if !(c.Request.Header.Get("Token")=="auth"){
			c.AbortWithStatusJSON(500,gin.H{
				"Message":"Token not present",
			})
			return
		}
		c.Next()
	}

} */