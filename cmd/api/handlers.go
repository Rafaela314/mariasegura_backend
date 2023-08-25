package main

import (
	"encoding/json"
	"fmt"
	"mariasegura/cmd/api/internal/models"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Maria running",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllVideos(w http.ResponseWriter, r *http.Request) {
	var videos []models.Video

	rd, _ := time.Parse("2006-01-02", "1986-03-07")

	highlander := models.Video{
		ID:          1,
		Title:       "highlander",
		ReleaseDate: rd,
		RunTime:     116,
		Description: "cool video",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	videos = append(videos, highlander)

	highlander2 := models.Video{
		ID:          2,
		Title:       "highlander",
		ReleaseDate: rd,
		RunTime:     116,
		Description: "cool video",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	videos = append(videos, highlander2)

	out, err := json.Marshal(videos)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
