// Matrikelnummern: 5911189 und 8441837
package handlers

import (
	"html/template"
	"meinProjekt/userManagement"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./templates/register.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		// Daten aus dem Formular abrufen
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Registrierung durchführen
		_, err := userManagement.Register(username, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Erfolgreich registriert, weiterleiten oder bestätigen
		http.Redirect(w, r, "/success", http.StatusSeeOther)
	}
}
