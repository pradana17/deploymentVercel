package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

// User struct untuk data dummy
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Jane Doe"},
}

func main() {
	r := gin.Default()

	r.GET("/api/users", func(c *gin.Context) {
		c.JSON(200, users)
	})

	r.POST("/api/users", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		newUser.ID = len(users) + 1
		users = append(users, newUser)
		c.JSON(201, newUser)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	r.Run(":" + port)
}
