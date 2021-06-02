package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Encrypto07/Booking-App/pkg/config"
	"github.com/Encrypto07/Booking-App/pkg/handlers"
	"github.com/Encrypto07/Booking-App/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting Application on Port Number %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
