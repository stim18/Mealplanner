// Matrikelnummern: 5911189 und 8441837
package handlers

import (
	"fmt"
	"html/template"
	"meinProjekt/data"
	"meinProjekt/userManagement"
	"net/http"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./templates/login.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		// Daten aus dem Formular abrufen
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Registrierung durchführen
		_, err := userManagement.Login(username, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Erfolgreich registriert, weiterleiten oder bestätigen
		http.Redirect(w, r, "/currentuser", http.StatusSeeOther)
		//StartLogoutTimer(w, r)
	}
}

func StartLogoutTimer(w http.ResponseWriter, r *http.Request) {
	timer := time.NewTimer(10 * time.Second)
	<-timer.C
	logoutHandler(w, r)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	data.ClearCurrentUser() // Lösche den angemeldeten Benutzer
	fmt.Println("Automatischer Logout durchgeführt")

	// Weiterleitung zur Login-Seite nach dem Logout
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
