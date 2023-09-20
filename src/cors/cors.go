package cors

import (
	"github.com/gin-gonic/gin"
)

func Default() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Acess-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Acess-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Acess-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accep-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Acess-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
