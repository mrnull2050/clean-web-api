package main

import (
	"net/http"

	swaggerfile "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{
		v1.GET("/eventes", app.GetEvents)
		v1.GET("/eventes/:id", app.GetEvent)

		v1.GET("/event/:id/attandees/", app.GetAttandeesForEvent)
		v1.GET("/attendee/:id/events", app.GetEventByAttendees)

		v1.POST("/auth/register", app.RegisterUser)
		v1.POST("/user/login", app.login)
	}
	authGroup := v1.Group("/")
	authGroup.Use(app.AuthMiddleWare())
	{
		authGroup.PUT("/event/:id", app.updateEvent)
		authGroup.DELETE("/event/:id", app.DeleteEvent)
		authGroup.POST("/events/:id/attendees/:userId", app.addAttendeestoEvent)
		authGroup.POST("/eventes", app.CreateEvent)
		authGroup.DELETE("/event/:id/attendees/:userId", app.DeleteAttendeeFromEvent)

	}

	g.GET("/swagger/*any", func(ctx *gin.Context) {
		if ctx.Request.RequestURI == "/swagger/" {
			ctx.Redirect(302, "/swagger/index.html")

		}
		ginswagger.WrapHandler(swaggerfile.Handler, ginswagger.URL("http://localhost:8080/swagger/doc.json"))(ctx)

	})

	return g
}
