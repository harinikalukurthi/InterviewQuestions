package main

import (
	"math/rand"
	"net/http"
	"sort"
	"sync"

	"github.com/gin-gonic/gin"
)

var urlStore = make(map[string]string)
var reverseURLStore = make(map[string]string)

var domainCount = make(map[string]int)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var mu sync.Mutex

func main() {
	router := gin.Default()
	router.POST("/shorten", shortenURL)
	router.GET("/:shortURL", redirectURL)
	router.GET("/metrics", getTopDomains)
	router.DELETE("/delete/:shortURL", deleteUrL)
	router.GET("/displayURL", listURL)
	router.GET("/list/:shortURL", listURLbyShortURL)
	router.GET("/count/:shortURL",listNoOfTimes)
	router.Run("localhost:8080")
}

func generateShortURL(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func shortenURL(c *gin.Context) {
	var request struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mu.Lock()
	defer mu.Unlock()

	if shortURL, exists := reverseURLStore[request.URL]; exists {
		c.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:8080/" + shortURL})
		domainCount[shortURL]++
		return
	}

	shortURL := generateShortURL(6)
	urlStore[shortURL] = request.URL
	reverseURLStore[request.URL] = shortURL
	domainCount[shortURL]++

	c.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:8080/" + shortURL})
}

func redirectURL(c *gin.Context) {
	shortURL := c.Param("shortURL")

	mu.Lock()
	originalURL, exists := urlStore[shortURL]
	mu.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}

func getTopDomains(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	type domainCountPair struct {
		Domain string
		Count  int
	}
	counts := make([]domainCountPair, 0, len(domainCount))
	for domain, count := range domainCount {
		counts = append(counts, domainCountPair{Domain: domain, Count: count})
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].Count > counts[j].Count
	})

	if len(counts) > 3 {
		counts = counts[:3]
	}

	c.JSON(http.StatusOK, gin.H{"top_domains": counts})
}

func deleteUrL(c *gin.Context) {
	shortURL := c.Param("shortURL")

	mu.Lock()
	defer mu.Unlock()
	if originalURL, exists := urlStore[shortURL]; exists {
		delete(urlStore, shortURL)
		delete(reverseURLStore, originalURL)
		delete(domainCount,shortURL)
		c.JSON(http.StatusOK, urlStore)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "url not found"})

}

func listURL(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	c.JSON(http.StatusOK, urlStore)
}

func listURLbyShortURL(c *gin.Context) {
	url := c.Param("shortURL")
	mu.Lock()
	defer mu.Unlock()

	if originalURL, exists := urlStore[url]; exists {
		c.JSON(http.StatusOK, originalURL)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "url not found"})
}

func listNoOfTimes(c *gin.Context){
	url:=c.Param("shortURL")

	mu.Lock()
	defer mu.Unlock()

	if count,exists:=domainCount[url];exists{
        c.JSON(http.StatusOK,count)
		return
	}
	c.JSON(http.StatusNotFound,gin.H{"message":"not found"})
}