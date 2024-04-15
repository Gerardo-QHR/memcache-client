package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gin-gonic/gin"
)

var memcachedClient *memcache.Client

func main() {
	// Initialize Memcached client
	memcachedClient = memcache.New("memcached:11211")
	memcachedClient.Timeout = 5 * time.Second

	// Initialize Gin router
	router := gin.Default()

	// Define routes
	router.POST("/set/:key", setKey)
	router.GET("/get/:key", getKey)

	// Run the server
	router.Run(":8080")
}

// Handler for setting key-value pair in Memcached
func setKey(c *gin.Context) {
	key := c.Param("key")
	value := c.Query("value")

	// Set expiration time to 3600 seconds (1 hour)
	expiration := 3600

	// Set key-value pair in Memcached with the specified expiration time
	err := memcachedClient.Set(&memcache.Item{Key: key, Value: []byte(value), Expiration: int32(expiration)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to set key '%s' in Memcached: %v", key, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Key '%s' set successfully in Memcached with a TTL of 3600 seconds", key)})
}

// Handler for getting value from Memcached
func getKey(c *gin.Context) {
	key := c.Param("key")

	// Get value from Memcached
	item, err := memcachedClient.Get(key)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Key '%s' not found", key)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get value for key '%s' from Memcached", key)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"value": string(item.Value)})
}
