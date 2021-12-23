package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a router
	r := gin.Default()

	// GET takes a route and a handler function
	// handler takes the gin context object
	r.GET("/pingTime", func(c *gin.Context) {
		// JSON serializer is available on gin context
		// Serialize to JSON data to client
		c.JSON(200, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})

	// Listen on port 8000
	r.Run(":8000")
}
