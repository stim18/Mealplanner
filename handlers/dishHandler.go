// Matrikelnummern: 5911189 und 8441837
package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"meinProjekt/content"
	"meinProjekt/data"
	"net/http"
)

func DishHandler(w http.ResponseWriter, r *http.Request) {
	currentUser := data.GetCurrentUser()

	tmpl := template.Must(template.ParseFiles("./templates/dish.html"))
	tmpl.Execute(w, currentUser)
}

func GetDish(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(content.DM.DishMap)
	if err != nil {
		http.Error(w, "Error encoding dishes", http.StatusInternalServerError)
		return
	}
}

func DeleteDishHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Fehler beim Lesen des Anfragekörpers", http.StatusBadRequest)
		return
	}

	var dishData map[string]interface{}
	err = json.Unmarshal([]byte(body), &dishData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	dishName := fmt.Sprintf("%v", dishData["Name"])

	err = content.DM.DeleteDish(dishName)
	if err != nil {
		fmt.Println("Error deleting dish:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Daten erfolgreich empfangen und verarbeitet!")
}

func AddDishHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Fehler beim Lesen des Anfragekörpers", http.StatusBadRequest)
		return
	}

	var groceryData map[string]interface{}
	err = json.Unmarshal(body, &groceryData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	var dishData struct {
		Name            string `json:"DishName"`
		DishComposition []struct {
			DishGroceries    string  `json:"DishGroceries,omitempty"`
			DishDishes       string  `json:"DishDishes,omitempty"`
			DishQuantityGram float64 `json:"DishQuantityGram"`
			DishQuantityUnit float64 `json:"DishQuantityUnit"`
		} `json:"DishComposition"`
	}

	dishName, ok := groceryData["DishName"].(string)
	if !ok || dishName == "" {
		http.Error(w, "Ungültige Eingabe: Gerichtsname oder Komposition fehlen", http.StatusBadRequest)
		return
	}
	dishData.Name = dishName

	composition, ok := groceryData["DishComposition"].([]interface{})
	if !ok || len(composition) == 0 {
		http.Error(w, "Ungültige Eingabe: Gerichtsname oder Komposition fehlen", http.StatusBadRequest)
		return
	}

	for _, comp := range composition {
		item, ok := comp.(map[string]interface{})
		if !ok {
			continue
		}

		var dishComp struct {
			DishGroceries    string  `json:"DishGroceries,omitempty"`
			DishDishes       string  `json:"DishDishes,omitempty"`
			DishQuantityGram float64 `json:"DishQuantityGram"`
			DishQuantityUnit float64 `json:"DishQuantityUnit"`
		}

		if groceries, ok := item["DishGroceries"].(string); ok {
			dishComp.DishGroceries = groceries
		}
		if dishes, ok := item["DishDishes"].(string); ok {
			dishComp.DishDishes = dishes
		}
		if quantityGram, ok := item["DishQuantityGram"].(float64); ok {
			dishComp.DishQuantityGram = quantityGram
		}
		if quantityUnit, ok := item["DishQuantityUnit"].(float64); ok {
			dishComp.DishQuantityUnit = quantityUnit
		}

		dishData.DishComposition = append(dishData.DishComposition, dishComp)
	}

	if dishData.Name == "" || len(dishData.DishComposition) == 0 {
		http.Error(w, "Ungültige Eingabe: Gerichtsname oder Komposition fehlen", http.StatusBadRequest)
		return
	}

	var newContents []content.DishContents
	for _, dishContent := range dishData.DishComposition {
		if dishContent.DishGroceries != "" {
			groceriesName := dishContent.DishGroceries
			quantityGram := dishContent.DishQuantityGram
			quantityUnit := dishContent.DishQuantityUnit

			groceries, exists := content.GM.GroceriesMap[groceriesName]
			if !exists {
				http.Error(w, "Ungültiges Lebensmittel in der Komposition", http.StatusBadRequest)
				return
			}

			newContents = append(newContents, content.DishContents{
				DishGroceries:    &groceries,
				DishQuantityGram: quantityGram,
				DishQuantityUnit: quantityUnit,
			})
		} else if dishContent.DishDishes != "" {
			dishName := dishContent.DishDishes
			quantityGram := dishContent.DishQuantityGram
			quantityUnit := dishContent.DishQuantityUnit

			dish, exists := content.DM.DishMap[dishName]
			if !exists {
				http.Error(w, "Ungültiges Gericht in der Komposition", http.StatusBadRequest)
				return
			}

			newContents = append(newContents, content.DishContents{
				DishDishes:       &dish,
				DishQuantityGram: quantityGram,
				DishQuantityUnit: quantityUnit,
			})
		}
	}
	// Annahme: content.DM repräsentiert deine DishManager-Instanz
	err = content.DM.NewDish(dishData.Name, newContents)
	if err != nil {
		fmt.Println("Error adding dish:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Daten erfolgreich empfangen und verarbeitet!")
}

func EditDishHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Fehler beim Lesen des Anfragekörpers", http.StatusBadRequest)
		return
	}

	var groceryData map[string]interface{}
	err = json.Unmarshal(body, &groceryData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	var dishData struct {
		Name            string `json:"DishName"`
		DishComposition []struct {
			DishGroceries    string  `json:"DishGroceries,omitempty"`
			DishDishes       string  `json:"DishDishes,omitempty"`
			DishQuantityGram float64 `json:"DishQuantityGram"`
			DishQuantityUnit float64 `json:"DishQuantityUnit"`
		} `json:"DishComposition"`
	}

	dishName, ok := groceryData["DishName"].(string)
	if !ok || dishName == "" {
		http.Error(w, "Ungültige Eingabe: Gerichtsname oder Komposition fehlen", http.StatusBadRequest)
		return
	}
	dishData.Name = dishName

	composition, ok := groceryData["DishComposition"].([]interface{})
	if !ok || len(composition) == 0 {
		http.Error(w, "Ungültige Eingabe: Gerichtsname oder Komposition fehlen", http.StatusBadRequest)
		return
	}

	for _, comp := range composition {
		item, ok := comp.(map[string]interface{})
		if !ok {
			continue
		}

		var dishComp struct {
			DishGroceries    string  `json:"DishGroceries,omitempty"`
			DishDishes       string  `json:"DishDishes,omitempty"`
			DishQuantityGram float64 `json:"DishQuantityGram"`
			DishQuantityUnit float64 `json:"DishQuantityUnit"`
		}

		if groceries, ok := item["DishGroceries"].(string); ok {
			dishComp.DishGroceries = groceries
		}
		if dishes, ok := item["DishDishes"].(string); ok {
			dishComp.DishDishes = dishes
		}
		if quantityGram, ok := item["DishQuantityGram"].(float64); ok {
			dishComp.DishQuantityGram = quantityGram
		}
		if quantityUnit, ok := item["DishQuantityUnit"].(float64); ok {
			dishComp.DishQuantityUnit = quantityUnit
		}

		dishData.DishComposition = append(dishData.DishComposition, dishComp)
	}

	if dishData.Name == "" || len(dishData.DishComposition) == 0 {
		http.Error(w, "Ungültige Eingabe: Gerichtsname oder Komposition fehlen", http.StatusBadRequest)
		return
	}

	var newContents []content.DishContents
	for _, dishContent := range dishData.DishComposition {
		if dishContent.DishGroceries != "" {
			groceriesName := dishContent.DishGroceries
			quantityGram := dishContent.DishQuantityGram
			quantityUnit := dishContent.DishQuantityUnit

			groceries, exists := content.GM.GroceriesMap[groceriesName]
			if !exists {
				http.Error(w, "Ungültiges Lebensmittel in der Komposition", http.StatusBadRequest)
				return
			}

			newContents = append(newContents, content.DishContents{
				DishGroceries:    &groceries,
				DishQuantityGram: quantityGram,
				DishQuantityUnit: quantityUnit,
			})
		} else if dishContent.DishDishes != "" {
			dishName := dishContent.DishDishes
			quantityGram := dishContent.DishQuantityGram
			quantityUnit := dishContent.DishQuantityUnit

			dish, exists := content.DM.DishMap[dishName]
			if !exists {
				http.Error(w, "Ungültiges Gericht in der Komposition", http.StatusBadRequest)
				return
			}

			newContents = append(newContents, content.DishContents{
				DishDishes:       &dish,
				DishQuantityGram: quantityGram,
				DishQuantityUnit: quantityUnit,
			})
		}
	}
	// Annahme: content.DM repräsentiert deine DishManager-Instanz
	err = content.DM.EditDishAttributeContent(dishData.Name, newContents)
	if err != nil {
		fmt.Println("Error adding dish:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Daten erfolgreich empfangen und verarbeitet!")
}
