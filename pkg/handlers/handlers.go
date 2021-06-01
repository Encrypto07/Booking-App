package handlers

import (
	"net/http"

	"github.com/Encrypto07/Booking-App/pkg/config"
	"github.com/Encrypto07/Booking-App/pkg/models"
	"github.com/Encrypto07/Booking-App/pkg/render"
)

//Repo the repository used by the handler
var Repo *Repository

//Repositery is the repositery type
type Repository struct {
	App *config.AppConfig
}

//NewrRepo creats a new repositery
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap})

}
