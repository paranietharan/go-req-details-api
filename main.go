package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		method := c.Request.Method
		path := c.Request.URL.Path
		status := c.Writer.Status()

		c.JSON(
			status,
			gin.H{
				"clientIP":  clientIP,
				"userAgent": userAgent,
				"method":    method,
				"path":      path,
			},
		)
	})

	r.Run(":8080") // localhost:8080
}
