package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	validUsers := map[string]string{
		"dev": "dev",
	}

	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username dan password diperlukan"})

			c.Abort()
			return
		}

		if validPassword, found := validUsers[username]; !found || validPassword != password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau password salah"})

			c.Abort()
			return
		}

		c.Next()
	}
}
