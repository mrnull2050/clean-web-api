package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrnull2050/clean-web-api/internal/database"
)

func (app *application) CreateEvent(c *gin.Context) {
	var Event database.Event

	if err := c.ShouldBindJSON(Event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := app.models.Event.Insert(&Event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can not insert data to evet in Create event "})
		return

	}
	c.JSON(http.StatusCreated, Event)

}

func (app *application) GetEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return

	}
	event, err := app.models.Event.Get(id)
	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found this event"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}

func (app *application) updateEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ExistingEvent, err := app.models.Event.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error : ": "could not get Event"})
		return
	}
	if ExistingEvent == nil {
		c.JSON(http.StatusNotFound, gin.H{"error : ": "NOT found Event"})
		return
	}
	UpdateEvent := app.models.Event.Update(ExistingEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, UpdateEvent)
}

func (app *application) DeleteEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	event := app.models.Event.Delete(id)
	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found Event"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "delete susseccfuly")
}
