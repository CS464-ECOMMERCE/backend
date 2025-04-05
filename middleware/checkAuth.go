package middleware

import (
	"backend/configs"
	"backend/models"
	"fmt"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CheckAuth(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	fmt.Println("1", tokenString)
	if err != nil {
		fmt.Println("2", err)
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
		fmt.Println("3", err, !token.Valid)
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("4", ok, claims)
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	if !claims.VerifyExpiresAt(jwt.TimeFunc().Unix(), true) {
		fmt.Println("5", claims.VerifyExpiresAt(jwt.TimeFunc().Unix(), true))
		c.JSON(401, gin.H{"error": "Unauthorized token expired"})
		c.Abort()
		return
	}

	userRole := claims["role"].(string)

	authenticatedRoles := []string{string(models.RoleAdmin), string(models.RoleMerchant)}
	if !slices.Contains(authenticatedRoles, userRole) {
		fmt.Println("6", userRole)
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Set("user_id", claims["user_id"])
	c.Set("role", userRole)
	c.Next()
}
