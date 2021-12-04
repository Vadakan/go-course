package Render

import (
	"bytes"
	"github.com/Vadakan/go-course/pkg/Handlers"
	"github.com/Vadakan/go-course/pkg/config"

	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var (
	functions = make(template.FuncMap)
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string,td *Handlers.TemplateData) {

	tc := make(map[string]*template.Template)

	if app.UseCache == true {

		tc = app.TemplateCache
	} else {

		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get the template..")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func CreateTemplateCache() (mych map[string]*template.Template, err error) {

	mycache := make(map[string]*template.Template)

	pages, err := filepath.Glob("C:/Users/91910/Documents/goworkspace/src/github.com/Webapp-New/templates/*.page.html")

	if err != nil {
		return mycache, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return mycache, err
		}
		matches, err := filepath.Glob("C:/Users/91910/Documents/goworkspace/src/github.com/Webapp-New/templates/*.layout.html")
		if err != nil {
			return mycache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("C:/Users/91910/Documents/goworkspace/src/github.com/Webapp-New/templates/*.layout.html")
			if err != nil {
				return mycache, err
			}
		}
		mycache[name] = ts
	}
	return mycache, nil
}
