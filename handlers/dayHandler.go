// Matrikelnummern: 5911189 und 8441837
package handlers

import (
	"encoding/json"
	"html/template"
	"meinProjekt/content"
	"meinProjekt/data"
	"net/http"
)

func DayHandler(w http.ResponseWriter, r *http.Request) {
	currentUser := data.GetCurrentUser()

	tmpl := template.Must(template.ParseFiles("./templates/day.html"))
	tmpl.Execute(w, currentUser)
}

func GetDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(content.DAYM.DayMap)
	if err != nil {
		http.Error(w, "Error encoding days", http.StatusInternalServerError)
		return
	}
}
