// Matrikelnummern: 5911189 und 8441837
package userManagement

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"meinProjekt/content"
	"meinProjekt/data"
	"os"
)

// SaveUserData speichert die Daten in einer Datei username.json im Ordner data
func SaveUserData(username string) error {
	// Daten in JSON umwandeln
	data := map[string]interface{}{
		"dayMap":       content.DAYM.DayMap,
		"mealMap":      content.MM.MealMap,
		"dishMap":      content.DM.DishMap,
		"groceriesMap": content.GM.GroceriesMap,
	}

	// Konvertiere Daten zu JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// Verschlüssle die Daten
	encryptedData, err := Encrypt(jsonData)
	if err != nil {
		return err
	}

	// Datei für die Daten erstellen
	filepath := fmt.Sprintf("data/%s.json", username)

	// Prüfen, ob die Datei bereits existiert
	if _, err := os.Stat(filepath); err == nil {
		// Datei existiert, daher löschen
		if err := os.Remove(filepath); err != nil {
			return err
		}
	}

	if err := ioutil.WriteFile(filepath, encryptedData, 0644); err != nil {
		return err
	}

	return nil
}

func Encrypt(dataToHash []byte) ([]byte, error) {
	currentUser := data.GetCurrentUser()
	passphrase := []byte(currentUser.Password)

	block, err := aes.NewCipher(passphrase)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(dataToHash))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], dataToHash)

	return ciphertext, nil
}
