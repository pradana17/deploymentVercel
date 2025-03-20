package handler

import (
	"net/http"

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

func Handler(w http.ResponseWriter, r *http.Request) {
	rg := gin.Default()

	rg.GET("/api/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	rg.POST("/api/users", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newUser.ID = len(users) + 1
		users = append(users, newUser)
		c.JSON(http.StatusCreated, newUser)
	})

	rg.ServeHTTP(w, r)
}
