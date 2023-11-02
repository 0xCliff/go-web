package handlers

import (
	"net/http"

	"github.com/0xCliff/basic-web-app/pkg/config"
	"github.com/0xCliff/basic-web-app/pkg/models"
	"github.com/0xCliff/basic-web-app/pkg/render"
)

// Respository type
type Repository struct {
	App *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handler
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page router
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := map[string]string{}
	stringMap["test"] = "Hello, again!"

	// send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
