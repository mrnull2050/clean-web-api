package main

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mrnull2050/clean-web-api/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type Registerrequired struct {
	Email    string `json:"email"  binding:"required,email"`
	Password string `json:"password"  binding:"required,min=8"`
	Name     string `json:"name"  binding:"required,min=2"`
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
