package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	// HTTP Server
	r := gin.Default()
	r.GET("/api/products/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Product Service: Hello!"})
	})

	r.GET("/api/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		val, err := rdb.Get(ctx, "product:"+id).Result()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"product": val})
	})

	r.Run(":8080")
}
