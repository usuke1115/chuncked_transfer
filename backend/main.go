package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CorsMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Allow-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(CorsMiddle())

	router.GET("/stream", func(c *gin.Context) {
		c.Header("Content-Type", "text/plain")
		c.Header("Cache-Control", "no-cache")

		for i := 0; i < 10; i++ {
			fmt.Fprintf(c.Writer, "message %d\n", i)
			c.Writer.Flush()
			time.Sleep(time.Second)
		}
	})
	router.Run()
}
