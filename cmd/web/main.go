package main

import (
	"fmt"
	"github.com/Vadakan/go-course/pkg/Handlers"
	"github.com/Vadakan/go-course/pkg/Render"
	"github.com/Vadakan/go-course/pkg/config"

	"log"
	"net/http"
)

var (
	portNumber = ":8080"
)

// main is the main executable function
func main() {

	var app config.AppConfig

	Tc, err := Render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err.Error())
	}

	app.TemplateCache = Tc
	app.UseCache = false

	repo := Handlers.NewRepo(&app)

	Handlers.NewHandler(repo)

	Render.NewTemplate(&app)

	http.HandleFunc("/home", repo.Home)
	http.HandleFunc("/about", repo.About)

	fmt.Println(fmt.Sprintf("starting the application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
