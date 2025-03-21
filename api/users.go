package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router := gin.Default()

	router.GET("/api/users", func(c *gin.Context) {
		c.JSON(200, users)
	})

	router.POST("/api/users", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		newUser.ID = len(users) + 1
		users = append(users, newUser)
		c.JSON(201, newUser)
	})

	router.ServeHTTP(w, r)
}
