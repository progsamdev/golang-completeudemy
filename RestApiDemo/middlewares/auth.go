package middlewares

import (
	"net/http"
	"restapidemo/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {

	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Set("userID", userID)
	c.Next()
}
