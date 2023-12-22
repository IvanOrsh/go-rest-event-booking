package models

import "time"

type Event struct {
	ID          string
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() {
	// TODO: store in a db
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
