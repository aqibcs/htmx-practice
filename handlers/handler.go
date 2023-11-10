package handlers

import (
	"html/template"
	"htmx/models"
	"net/http"
	"time"
)

// Handler function #1 - returns the index.html template, with film data
func H1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	films := map[string][]models.Film{
		"Films": {
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "Blade Runner", Director: "Ridley Scott"},
			{Title: "The Thing", Director: "John Carpenter"},
		},
	}
	tmpl.Execute(w, films)
}

// Handler function #2 - returns the template block with the newly added film, as an HTMX response
func H2(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.ExecuteTemplate(w, "film-list-element", models.Film{Title: title, Director: director})
}
