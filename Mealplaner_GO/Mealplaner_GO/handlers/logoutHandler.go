// Matrikelnummern: 5911189 und 8441837
package handlers

import (
	"meinProjekt/data"
	"meinProjekt/userManagement"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	err := userManagement.SaveUserData(data.GetCurrentUser().Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.ClearCurrentUser()

	println("Current User: " + data.GetCurrentUser().Username)
}
