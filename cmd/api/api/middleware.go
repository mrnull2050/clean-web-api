package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (app *application) AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authhandler := ctx.GetHeader("Authorization")
		if authhandler == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invaildd"})
			ctx.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authhandler, "Bearer")
		if tokenString == authhandler {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invaildd {need Bearer}"})
			ctx.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(app.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invaild token"})
			ctx.Abort()
			return
		}
		claims , ok := token.Claims.(jwt.MapClaims)
		if !ok {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invaild token"})
			ctx.Abort()
			return
		}
	}
}
