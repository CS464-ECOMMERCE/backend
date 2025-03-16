package middleware

import (
	"backend/configs"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CheckAuth(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// Check if token is valid
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(configs.JWT_SECRET), nil
	})
	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	if !claims.VerifyExpiresAt(jwt.TimeFunc().Unix(), true) {
		c.JSON(401, gin.H{"error": "Unauthorized token expired"})
		c.Abort()
		return
	}

	c.Set("user_id", claims["user_id"])
	c.Set("role", claims["role"])
	c.Next()
}