package model

import "time"

// EventModel model
type EventModel struct {
	EventTime   time.Time `json:"event_time" gorm:"column:event_time;not null" binding:"required"`
	Country     string    `json:"country" gorm:"column:country"`
	Province    string    `json:"province" gorm:"column:province"`
	City        string    `json:"city" gorm:"column:city"`
	Description string    `json:"description" gorm:"column:description"`
}
