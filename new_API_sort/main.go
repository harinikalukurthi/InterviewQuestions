package main

import (
	"math/rand"
	"net/http"
	"sync"
	"sort"
	"github.com/gin-gonic/gin"
	"time"
)

type sample struct {
	originalURL string
	shortURL    string
	lastTime    time.Time
}

var (
	urlStore       = make(map[string]*sample)
	letters        = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	mu             sync.Mutex
	domainCount    = make(map[string]int)
)

func main() {
	router := gin.Default()
	router.POST("/shorten", shortenURL)
	router.GET("/:shortURL", redirectURL)
	router.GET("/metrics", getTopDomains)
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

	// Check if URL is already shortened
	for _, s := range urlStore {
		if s.originalURL == request.URL {
			s.lastTime = time.Now()
			c.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:8080/" + s.shortURL})
			return
		}
	}

	shortURL := generateShortURL(6)
	urlStore[shortURL] = &sample{
		originalURL: request.URL,
		shortURL:    shortURL,
		lastTime:    time.Now(),
	}

	domainCount[request.URL]++

	c.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:8080/" + shortURL})
}

func redirectURL(c *gin.Context) {
	shortURL := c.Param("shortURL")

	mu.Lock()
	s, exists := urlStore[shortURL]
	mu.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	if time.Since(s.lastTime) < 10*time.Second {
		s.lastTime = time.Now()
		c.Redirect(http.StatusMovedPermanently, s.originalURL)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "last accessed time is more than 10 seconds"})
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
