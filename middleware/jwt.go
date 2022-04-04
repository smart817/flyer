package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/smart817/flyer"
	"github.com/smart817/flyer/resp"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("x-token")
		if tokenHeader == "" {

			resp.Error(c, "no x-token")
			c.Abort()
			return
		}
		name, err := flyer.ParseToken(tokenHeader)
		if name == "" {
			resp.Error(c, err.Error()) //"token error"
		} else if err != nil {
			resp.Info(c, "token expire")
			c.Abort()
			return
		} else {
			c.Set("username", name)
			c.Next()
		}

	}
}
