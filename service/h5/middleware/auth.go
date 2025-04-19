package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		log.Printf("[%s] %s - %v", c.Request.Method, c.Request.URL.Path, latency)
	}
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatusJSON(500, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()
		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里可以添加认证逻辑
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		// 验证token逻辑...
		c.Next()
	}
}