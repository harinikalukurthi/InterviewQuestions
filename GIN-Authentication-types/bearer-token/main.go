package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var tokens []string

func main() {
	fmt.Println(tokens)
	router := gin.Default()
	router.GET("/token", generateToken)
	fmt.Println(tokens)
	router.GET("/data", getdata)
	router.Run()
}
func getdata(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	reqToken := strings.Split(bearerToken, " ")[1]
	for _, token := range tokens {
		if token == reqToken {
			c.JSON(http.StatusOK, gin.H{"message": "Hello world"})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "unauthorized",
	})
}

func generateToken(c *gin.Context) {
	token, _ := randomHex(20)
	tokens = append(tokens, token)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	fmt.Println(bytes)
	return hex.EncodeToString(bytes), nil
}


//curl -H "Authorization: Bearer 98aaa9385d265eacaf2ed3a947c13464ece16c81" -X GET http://localhost:8080/data