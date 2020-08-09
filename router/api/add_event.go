package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leetpy/daisy/lib/es"
)

type Event struct {
	Title       string    `form:"title"`
	Date        time.Time `form:"date"`
	Country     string    `form:"country"`
	Province    string    `form:"province"`
	City        string    `form:"city"`
	Longitude   float32   `form:"longitude"`
	Latitude    float32   `form:"latitude"`
	Description string    `form:"description"`
}

func AddEventHandler(c *gin.Context) {
	var event Event

	if c.ShouldBind(&event) != nil {
		c.JSON(http.StatusOK, gin.H{"error_message": "params are invalid!", "error_code": 1})
		return
	}

	// 丢入 es
	daisyEvent := es.Event{
		Timestamp:   time.Now(),
		Title:       event.Title,
		Date:        event.Date,
		Country:     event.Country,
		Province:    event.Province,
		City:        event.City,
		Location:    fmt.Sprintf("%f,%f", event.Latitude, event.Longitude),
		Description: event.Description,
	}
	err := daisyEvent.Create()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error_code":    1,
			"error_message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error_code":    0,
			"error_message": "success",
		})
	}
}
