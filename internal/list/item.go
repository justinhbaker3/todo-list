package list

import (
	"time"
)

// Item defines an Item in a list
type Item struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TimeOpened  time.Time `json:"time_opened"`
}

// New creates a new Item
func NewItem(title, description string) *Item {
	return &Item{
		Title:       title,
		Description: description,
		TimeOpened:  time.Now().UTC(),
		// completed and timeCompleted are automatically filled in with their zero values
	}
}
