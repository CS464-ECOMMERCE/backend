package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CheckSession(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")
	if err != nil || sessionID == "" {
		sessionID = uuid.New().String()
		c.SetCookie("session_id", sessionID, 60*60*24, "/", "", false, false)
	}

	c.Set("session_id", sessionID)
	c.Next()
}
