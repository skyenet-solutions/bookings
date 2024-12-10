package handlers

import (
	"net/http"

	"github.com/skyenet-solutions/bookings/pkg/config"
	"github.com/skyenet-solutions/bookings/pkg/models"
	"github.com/skyenet-solutions/bookings/pkg/render"
)

//Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new reopsitory
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request){

	remoteIP := r.RemoteAddr
	m.App.SessionManager.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request){

	//Perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.SessionManager.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	//Send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}

