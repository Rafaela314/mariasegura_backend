package models

import "time"

type Video struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	ReleaseDate  time.Time `json:"release_date"`
	RunTime      int       `json:"runtime"`
	Rating       float64   `json:"rating"`
	Description  string    `json:"description"`
	Image        string    `json:"image"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	ExternalLink string    `json:"external_link"`
}
