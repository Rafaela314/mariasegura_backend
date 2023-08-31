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

func (app *application) AllResources(w http.ResponseWriter, r *http.Request) {
	var resources []models.Resource

	rd, _ := time.Parse("2006-01-02", "1986-03-07")

	highlander := models.Resource{
		ID:             1,
		Title:          "highlander",
		RunTime:        116,
		Description:    "cool Ebook",
		Classification: "PDF",
		CreatedAt:      rd,
		UpdatedAt:      time.Now(),
	}

	resources = append(resources, highlander)

	highlander2 := models.Resource{
		ID:             2,
		Title:          "highlander",
		RunTime:        116,
		Description:    "cool video",
		Classification: "video",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	resources = append(resources, highlander2)

	out, err := json.Marshal(resources)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
