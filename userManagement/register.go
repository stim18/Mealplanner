// Matrikelnummern: 5911189 und 8441837
package userManagement

import (
	"encoding/csv"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"meinProjekt/data"
	"os"
	"regexp"
	"strings"
)

func Register(username, password string) (*data.User, error) {
	csvFilePath := "data/users.csv"

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

	// Überprüfen, ob die Datei existiert
	_, err := os.Stat(csvFilePath)
	isFileExist := !os.IsNotExist(err)

	// Passwort hashen
	hashedPassword := hashPassword(password)

	// Neue Benutzerdaten
	newUserData := []string{username, hashedPassword}

	// Öffnen der CSV-Datei im Lese-Modus
	file, err := os.Open(csvFilePath)
	if err != nil {
		if os.IsNotExist(err) && !isFileExist {
			// Erstelle die Datei, wenn sie nicht existiert
			file, err = os.Create(csvFilePath)
			if err != nil {
				return nil, err
			}
			defer file.Close()

			// Schreibe den Header in die neue Datei
			csvWriter := csv.NewWriter(file)
			defer csvWriter.Flush()

			err := csvWriter.Write([]string{"Username", "Password"})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		file.Close()
	}

	// Überprüfen, ob der Benutzer bereits existiert
	users, err := readCSV(csvFilePath)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user[0] == username {
			return nil, errors.New("Benutzer existiert bereits")
		}
	}

	// Öffnen der CSV-Datei im Schreibmodus, um den neuen Benutzer hinzuzufügen
	file, err = os.OpenFile(csvFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	// Schreibe die neuen Benutzerdaten in die CSV-Datei
	err = csvWriter.Write(newUserData)
	if err != nil {
		return nil, err
	}

	return &data.User{Username: username, Password: hashedPassword}, nil
}

func readCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func hashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
