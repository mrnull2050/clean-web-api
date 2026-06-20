package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {	
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{
		v1.POST("/eventes", app.CreateEvent)
		v1.GET("/eventes", app.GetEvent)
		v1.GET("/eventes/:id", app.GetEvent)
		v1.PUT("/event/:id", app.updateEvent)
		v1.DELETE("/event/:id", app.DeleteEvent)

		v1.POST("/auth/register", app.RegisterUser)

	}
	return g
}
