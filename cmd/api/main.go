package main

import (
	"log"
	"net/http"
)

type application struct {
	Domain string
}

func main() {

	var app application

	app.Domain = "mariasegura.com"

	err := http.ListenAndServe(":8080", app.routes())
	if err != nil {
		log.Fatal(err)
	}

}
