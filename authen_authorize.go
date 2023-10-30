// routes/middleware.go

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is used to protect routes that require authentication.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user is authenticated (e.g., by verifying session or tokens)
		// If not authenticated, redirect to login or return an unauthorized response
		// You can also set user data in the context for controllers to access
		if !userAuthenticated {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// AdminMiddleware is used to protect routes that require admin privileges.
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user has admin privileges
		// If not an admin, return an unauthorized response or redirect
		if !userIsAdmin {
			c.JSON(http.StatusForbidden, gin.H{"message": "Admin access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
