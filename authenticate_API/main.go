package main

import (
	"net/http"

	"github.com/go-gonic/gin"
)

func main() {
	type Login struct {
		User     string `form:"user" json:"user" xml:"user"  binding:"required"`
		Password string `form:"password" json:"password" xml:"password" binding:"required"`
	}

	router := gin.Default()

	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
}
