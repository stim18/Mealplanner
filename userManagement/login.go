// Matrikelnummern: 5911189 und 8441837
package userManagement

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"meinProjekt/data"
	"regexp"
	"strings"
)

func Login(username, password string) (*data.User, error) {
	csvFilePath := "data/users.csv"
	// Lesen der Benutzerdaten aus der CSV-Datei
	users, err := readCSV(csvFilePath)
	if err != nil {
		return nil, err
	}

	username = strings.ToLower(username)
	username = strings.TrimSpace(username)

	reUsername := regexp.MustCompile(`^[a-zA-Z0-9@.+\-_~]+$`)
	rePassword := regexp.MustCompile(`^[a-zA-Z0-9~!@#$%^&*()_\-+=<>?/{}[\]|;:',.]+$`)

	validUsername := reUsername.MatchString(username)
	validPassword := rePassword.MatchString(password)

	if !validUsername {
		return nil, errors.New("keine valide Benutzereingabe")
	}
	if !validPassword {
		return nil, errors.New("keine valide Passwordeingabe")
	}

	// Überprüfen, ob der Benutzer existiert und das Passwort korrekt ist
	for _, user := range users[0:] { // Überspringen des Header-Eintrags
		if user[0] == username {
			err := bcrypt.CompareHashAndPassword([]byte(user[1]), []byte(password))
			if err != nil {
				return nil, errors.New("Ungültiges Passwort")
			}
			// Speichern des angemeldeten Benutzers in data.CurrentUser
			data.SetCurrentUser(username, password) // Setze den aktuellen Benutzer
			err = LoadUserData(username)
			if err != nil {
				return nil, err
			}
			return &data.User{Username: username, Password: user[1]}, nil
		}
	}

	return nil, errors.New("Benutzer existiert nicht")
}
