package es

import (
	"context"
	"time"

	"github.com/leetpy/daisy/lib/utils"
)

type Event struct {
	Timestamp   time.Time `json:"@timestamp"`
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Country     string    `json:"country"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
}

func (e *Event) Create() error {
	index := utils.GetEventsIndexName()
	client := GetClient()
	_, err := client.Index().Index(index).Type("default").BodyJson(&e).Do(context.Background())
	return err
}
