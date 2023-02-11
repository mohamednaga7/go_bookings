package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/mohamednaga7/go_bookings/pkg/config"
	"github.com/mohamednaga7/go_bookings/pkg/handlers"
	"github.com/mohamednaga7/go_bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

var portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println("Starting web server on port 8080")

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
