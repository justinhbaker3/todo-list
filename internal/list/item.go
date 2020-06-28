package list

import (
	"time"
)

// item defines an item in a list
type item struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TimeOpened  time.Time `json:"time_opened"`
}

// New creates a new item
func NewItem(title, description string) *item {
	return &item{
		Title:       title,
		Description: description,
		TimeOpened:  time.Now().UTC(),
		// completed and timeCompleted are automatically filled in with their zero values
	}
}
