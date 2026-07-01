package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrnull2050/clean-web-api/internal/database"
)

func (app *application) GetUserFromContext(c *gin.Context) *database.User {
	userContext, exiest := c.Get("user")
	if !exiest {
		return &database.User{}

	}
	user, ok := userContext.(*database.User)
	if !ok {
		return &database.User{}
	}
	return user

}
