package routes

import (
	"log"
	"net/http"
	"restapidemo/Models"
	"restapidemo/utils"

	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {

	var user Models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := user.Save()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func login(c *gin.Context) {
	var user Models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := user.ValidateCredentials(user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not authenticate user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not authenticate user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
