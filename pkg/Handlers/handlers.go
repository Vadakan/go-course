package Handlers

import (
	"github.com/Vadakan/go-course/pkg/Render"
	"github.com/Vadakan/go-course/pkg/config"
	"net/http"
)

// TemplateData struct to pass data from handler to render
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

var (
	app *Repository
)

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	app = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	Render.RenderTemplate(w, "home.page.html",&TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some business logic
	StringMap := make(map[string]string)
	StringMap["test"] = "hello,again"

	//send the data to the template
	Render.RenderTemplate(w, "about.page.html",&TemplateData{
		StringMap: StringMap,
	})
}
