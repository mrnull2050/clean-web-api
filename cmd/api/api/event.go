package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrnull2050/clean-web-api/internal/database"
)

func (app *application) CreateEvent(c *gin.Context) {
	var Event database.Event

	if err := c.ShouldBindJSON(&Event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := app.GetUserFromContext(c)
	Event.OwnerId = user.Id
	err := app.models.Event.Insert(&Event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can not insert data to evet in Create event "})
		return

	}
	c.JSON(http.StatusCreated, Event)

}

// @Summary Returns all events
// @Description Returns all events
// @Tags Events
// @Accept json
// @Produce json
// @Success 200 {array} database.Event
// @Router /api/v1/events [get]
func (app *application) GetEvents(c *gin.Context) {
	evets, err := app.models.Event.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, evets)
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
		c.JSON(http.StatusBadRequest, gin.H{"error get id": err.Error()})
		return
	}
	ExistingEvent, err := app.models.Event.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error : ": "could not get Event"})
		return
	}
	user := app.GetUserFromContext(c)
	if ExistingEvent == nil {
		c.JSON(http.StatusNotFound, gin.H{"error : ": "NOT found Event"})
		return
	}
	if ExistingEvent.OwnerId != user.Id {
		c.JSON(403, gin.H{"error": "you are not authroized "})
		return
	}
	var input database.Event
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error input : ": err.Error()})
		return
	}
	ExistingEvent.Name = input.Name
	ExistingEvent.Description = input.Description
	ExistingEvent.Date = input.Date
	ExistingEvent.Location = input.Location

	UpdateEvent, err := app.models.Event.Update(ExistingEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error update": err.Error()})
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

	err = app.models.Event.Delete(id)
	if err == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found Event"})
		return
	}
	event, err := app.models.Event.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found the event"})
		return
	}
	user := app.GetUserFromContext(c)
	if event.OwnerId != user.Id {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not authorized to delete this event!!!"})
		return
	}

	c.JSON(http.StatusOK, "delete susseccfuly")
}

func (app *application) addAttendeestoEvent(c *gin.Context) {
	eventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "faild to Get Event ID"})
		return
	}
	UserID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "faild to get User ID"})
		return
	}
	event, err := app.models.Event.Get(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "faild to retive event"})
		return
	}
	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	user, err := app.models.Event.Get(UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "faild to retive user"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	existingAttendee, err := app.models.Attendees.GetByEventAndAttendee(eventId, UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "faild to retive attendees"})
		return
	}
	if existingAttendee != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "attendees already exist"})
		return
	}
	attendee := database.Attendees{
		EventId: eventId,
		UserId:  user.Id,
	}
	_, err = app.models.Attendees.Instert(&attendee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "faild to add attendees"})
		return
	}
	c.JSON(http.StatusCreated, attendee)
}

func (app *application) GetAttandeesForEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not get ID for Event"})
		return
	}
	user, err := app.models.Attendees.GetAttendeesByEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "faild to retive attendees for event"})
		return
	}
	c.JSON(http.StatusOK, user)
}
func (app *application) DeleteAttendeeFromEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not get ID for Event"})
		return
	}
	userid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not get ID for user"})
		return
	}
	event, err := app.models.Event.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found!!!"})
		return
	}
	user := app.GetUserFromContext(c)
	if event.OwnerId != user.Id {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not allow to delete attendee from event"})
	}
	err = app.models.Attendees.Delete(userid, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can not delete from db!!!"})
		return
	}
	c.JSON(http.StatusNoContent, nil)

}

func (app *application) GetEventByAttendees(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": "user id is not ok"})
		return
	}
	event, err := app.models.Attendees.GetByAttendees(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "con not Get event by attendees"})
	}
	c.JSON(http.StatusOK, event)
}
