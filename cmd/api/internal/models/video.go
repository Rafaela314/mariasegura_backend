package models

import "time"

type Resource struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	RunTime        int       `json:"runtime"`
	Description    string    `json:"description"`
	Image          string    `json:"image"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	ExternalLink   string    `json:"external_link"`
	Classification string    `json:"classification"`
}
