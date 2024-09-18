package middleware
import (
	"github.com/gin-gonic/gin"
)
func Authenticator(c *gin.Context){
	if !(c.Request.Header.Get("Token")=="auth"){
		c.AbortWithStatusJSON(500,gin.H{
			"Message":"Token not present",
		})
		return
	}
	c.Next()
}