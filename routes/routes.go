package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// THIS MOTHERFUCKING THING IS FUCKING CASE SENSITIVE, I'M ABOUT TO SCREAM
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User

func SetupRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	r.GET("/greetings/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s!", name)
	})

	r.POST("/users", func(c *gin.Context) {

		var newUser1 User
		if err := c.BindJSON(&newUser1); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
			return
		}
		fmt.Println(newUser1)
		if newUser1.Name == "" || newUser1.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name and email are required"})
			return
		}
		users = append(users, newUser1)

		c.JSON(http.StatusOK, gin.H{"message": "User added successfully", "users": users})
	})
}
