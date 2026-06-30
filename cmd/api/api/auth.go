package main

import (
	"net/http"
	"time"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mrnull2050/clean-web-api/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type Registerrequired struct {
	Email    string `json:"email"  binding:"required,email"`
	Password string `json:"password"  binding:"required,min=8"`
	Name     string `json:"name"  binding:"required,min=2"`
}

type LoginRequest struct {
	Email    string `json:"email"  binding:"required,email"`
	Password string `json:"password"  binding:"required,min=8"`
}
type LoginResponse struct {
	Token string `json:"password"  binding:"required,min=8"`
}

func (app *application) login(c *gin.Context) {
	var auth LoginRequest
	if err := c.ShouldBindJSON(auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	exiestingUser, err := app.models.User.GetByEmail(auth.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invaild Email or password"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "sth went wrong"})
	}
	err = bcrypt.CompareHashAndPassword([]byte(exiestingUser.Password), []byte(auth.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invaild Email or password"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": exiestingUser.Id,
		"expr":   time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(app.JWTSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, LoginResponse{Token: tokenString})
}
func (app *application) RegisterUser(c *gin.Context) {
	var register Registerrequired
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error : ": err.Error()})
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	register.Password = string(hashPassword)

	user := database.User{
		Email:    register.Email,
		Password: register.Password,
		Name:     register.Name,
	}
	err = app.models.User.Insert(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error : ": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
	body, _ := c.GetRawData()
	fmt.Println(string(body))

}
