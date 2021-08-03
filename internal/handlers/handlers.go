package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Encrypto07/Booking-App/internal/config"
	"github.com/Encrypto07/Booking-App/internal/forms"
	"github.com/Encrypto07/Booking-App/internal/models"
	"github.com/Encrypto07/Booking-App/internal/render"
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

//Home is the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

//About is about handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again"

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}

//Reservation reader the make a reservation page and display form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {

	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})

}

//PostReservation handles the posting of form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
	}
	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

//for Generals page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

//for majors page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

//avalability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

//postavalability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {

	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

//Avalibalityjson request for for handling the json
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      false,
		Message: "Available!",
	}
	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Panicln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

//Rendering The Contact page through handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("cannot get the item from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
