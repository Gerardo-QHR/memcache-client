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
	memcachedClient = memcache.New("memcached:11211")
	memcachedClient.Timeout = 5 * time.Second

	router := gin.Default()

	router.POST("/set/:key", setKey)
	router.GET("/get/:key", getKey)

	router.Run(":8080")
}

func setKey(c *gin.Context) {
	key := c.Param("key")
	value := c.Query("value")

	expiration := 3600

	err := memcachedClient.Set(&memcache.Item{Key: key, Value: []byte(value), Expiration: int32(expiration)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to set key '%s' in Memcached: %v", key, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Key '%s' set successfully in Memcached with a TTL of 3600 seconds", key)})
}

func getKey(c *gin.Context) {
	key := c.Param("key")

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
