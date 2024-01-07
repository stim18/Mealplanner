// Matrikelnummern: 5911189 und 8441837
package handlers

import (
	"html/template"
	"meinProjekt/data"
	"net/http"
)

func ShowCurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	// Holen des aktuellen Benutzers aus der Datenstruktur
	currentUser := data.GetCurrentUser()

	tmpl := template.Must(template.ParseFiles("./templates/start.html"))
	tmpl.Execute(w, currentUser)
}
