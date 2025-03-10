package main

import (
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Token")
		expectedToken := "Auth"
		// token = expectedToken
		if token != expectedToken {
			c.AbortWithStatusJSON(401, gin.H{
				"Message": "Invalid or missing token",
			})
			return
		}
		c.Next()
	}
}
func main() {
	router := gin.Default()
	router.Use(Authenticate())
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	router.Run()
}
