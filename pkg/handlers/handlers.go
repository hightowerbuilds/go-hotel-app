package handlers

import (
	"net/http"

	"github.com/hightowerbuilds/go-hotel-app/pkg/config"
	"github.com/hightowerbuilds/go-hotel-app/pkg/models"
	"github.com/hightowerbuilds/go-hotel-app/pkg/render"
)

// holds data sent from handlers to templates

// repo used by handlers
var Repo *Repository

// is reeposirtiy type
type Repository struct {
	App *config.AppConfig
}

// creates a new repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, from the inside"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
