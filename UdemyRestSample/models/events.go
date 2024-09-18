package models

import "time"

type Events struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Date        time.Time `json:"date"`
	UserID      int       `json:"userid"`
}

var Event = []Events{}

func (e Events) Save() {
	// adding into database
	Event = append(Event, e)
}

func GetAllEvents() []Events {
	return Event
}

