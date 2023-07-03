package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateSession(c *gin.Context, userID string,  userName string) {
	session := sessions.Default(c)
	session.Set("authenticated", true)
	session.Set("userID", userID)
	session.Set("userName", userName)
	session.Options(sessions.Options{
		Path:   "/",
		MaxAge: 600,
	})
	session.Save()
}

func IsAuthenticated(c *gin.Context) bool {
	session := sessions.Default(c)
	auth := session.Get("authenticated")
	return auth != nil && auth.(bool)
}

func GetUserID(c *gin.Context) string {
	session := sessions.Default(c)
	return session.Get("userID").(string)
}

func ClearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !IsAuthenticated(c) {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}