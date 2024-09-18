/*
endpoints list:

GET /events
GET /events/id
POST /events
PUT /events/id
DELETE /events/id
POST /signup
POST /login
POST /events/id/register
DELETE /events/id/register
*/
package main

import (
	"net/http"
	"strconv"

	"UdemyRestSample/models"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", addEvent)
	server.GET("/events/:id", getEventwithID)
	server.DELETE("/events/:id", deleteEvent)
	server.PUT("/events", updateEvent)

	server.Run()
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, events)
}

func addEvent(c *gin.Context) {
	var newEvent models.Events
	if err := c.BindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	newEvent.Save()
	c.JSON(http.StatusOK, newEvent)
}
func getEventwithID(c *gin.Context) {
	id := c.Param("id")
	id1, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "not able to convert to int"})
		return
	}
	for _, j := range models.Event {
		if j.Id == id1 {
			c.JSON(http.StatusOK, j)
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "did not find event", "event id": id})
}
func deleteEvent(c *gin.Context) {
	id := c.Param("id")
	id1, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "not able to convert to int"})
		return
	}
	for i, j := range models.Event {
		if j.Id == id1 {
			models.Event = append(models.Event[:i], models.Event[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"deleted event": j})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"could not find the event with id": id})
}

func updateEvent(c *gin.Context) {
	var updateEvent models.Events
	err := c.ShouldBindJSON(&updateEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cannot convert to json"})
		return
	}
	id := c.Query("id")
	id1, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "not able to convert to int"})
		return
	}

	for i, j := range models.Event {
		if j.Id == id1 {
			models.Event[i].Description = updateEvent.Description
			models.Event[i].Name = updateEvent.Name
			models.Event[i].Date = updateEvent.Date
			models.Event[i].Location = updateEvent.Location
			c.JSON(http.StatusOK, gin.H{"message": "updated", "event": updateEvent})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"could not find the event with id": id})
}
