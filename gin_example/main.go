package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve static HTML file
	r.StaticFile("/test", "static/index.html")

	// Create a custom router group
	api := r.Group("/api")
	{
		api.GET("/echo", getEchoHandler)
		api.GET("/echo/:id", getEchoWithIDHandler)
		api.POST("/echo", postEchoHandler)
	}

	r.Run() // Run on default port 8080
}

// getEchoHandler handles GET requests to echo a message via query parameters
func getEchoHandler(c *gin.Context) {
	message := c.Query("message")
	if message == "" {
		message = "No message provided"
	}
	c.JSON(http.StatusOK, gin.H{
		"echo": message,
	})
}

// getEchoWithIDHandler handles GET requests with an ID parameter
func getEchoWithIDHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID. ID must be an integer"})
		return
	}
	message := c.Query("message")
	if message == "" {
		message = "No message provided"
	}
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"echo": message,
	})
}

// postEchoHandler handles POST requests to echo a message from JSON body
func postEchoHandler(c *gin.Context) {
	var json struct {
		Message string `json:"message"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"echo": json.Message,
	})
}
