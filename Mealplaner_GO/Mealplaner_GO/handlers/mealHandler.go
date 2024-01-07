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

func MealHandler(w http.ResponseWriter, r *http.Request) {
	currentUser := data.GetCurrentUser()

	tmpl := template.Must(template.ParseFiles("./templates/meal.html"))
	tmpl.Execute(w, currentUser)
}

func GetMealHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(content.MM.MealMap)
	if err != nil {
		http.Error(w, "Error encoding meals", http.StatusInternalServerError)
		return
	}
}

func DeleteMealHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Fehler beim Lesen des Anfragekörpers", http.StatusBadRequest)
		return
	}

	var mealData map[string]interface{}
	err = json.Unmarshal([]byte(body), &mealData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	mealName := fmt.Sprintf("%v", mealData["Name"])

	err = content.MM.DeleteMeal(mealName)
	if err != nil {
		fmt.Println("Error deleting dish:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Daten erfolgreich empfangen und verarbeitet!")
}

func AddMealHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Fehler beim Lesen des Anfragekörpers", http.StatusBadRequest)
		return
	}

	var mealData struct {
		Name            string `json:"MealName"`
		MealComposition []struct {
			MealGroceries    string  `json:"MealGroceries,omitempty"`
			MealDishes       string  `json:"MealDishes,omitempty"`
			DishQuantityGram float64 `json:"DishQuantityGram"`
			DishQuantityUnit float64 `json:"DishQuantityUnit"`
		} `json:"MealComposition"`
	}

	err = json.Unmarshal(body, &mealData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	var newMealContent content.MealContents

	for _, mealContent := range mealData.MealComposition {
		if mealContent.MealGroceries != "" {
			groceriesName := mealContent.MealGroceries
			quantityGram := mealContent.DishQuantityGram
			quantityUnit := mealContent.DishQuantityUnit

			groceries, exists := content.GM.GroceriesMap[groceriesName]
			if exists {
				pointerToGroceries := &groceries

				newMealContent = content.MealContents{
					MealGroceries:    pointerToGroceries,
					DishQuantityGram: quantityGram,
					DishQuantityUnit: quantityUnit,
				}
			}
		} else if mealContent.MealDishes != "" {
			dishName := mealContent.MealDishes
			quantityGram := mealContent.DishQuantityGram
			quantityUnit := mealContent.DishQuantityUnit

			dish, exists := content.DM.DishMap[dishName]
			if exists {
				pointerToDish := &dish

				newMealContent = content.MealContents{
					MealDishes:       pointerToDish,
					DishQuantityGram: quantityGram,
					DishQuantityUnit: quantityUnit,
				}

			}
		}
	}

	var newMealData = content.Meal{
		Name:            mealData.Name,
		MealComposition: []content.MealContents{newMealContent},
	}

	err = content.MM.NewMeal(newMealData.Name, newMealData.MealComposition)
	if err != nil {
		fmt.Println("Error adding meal:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Daten erfolgreich empfangen und verarbeitet!")

	println("Meal added!")

	for key, value := range content.MM.MealMap {
		fmt.Println(key, ":", value)
	}
}

func EditMealHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Fehler beim Lesen des Anfragekörpers", http.StatusBadRequest)
		return
	}

	println(string(body))

	var mealData struct {
		Name            string `json:"MealName"`
		MealComposition []struct {
			MealGroceries    string  `json:"MealGroceries,omitempty"`
			MealDishes       string  `json:"MealDishes,omitempty"`
			DishQuantityGram float64 `json:"DishQuantityGram"`
			DishQuantityUnit float64 `json:"DishQuantityUnit"`
		} `json:"MealComposition"`
	}

	err = json.Unmarshal(body, &mealData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	var newMealContent content.MealContents

	for _, mealContent := range mealData.MealComposition {
		if mealContent.MealGroceries != "" {
			groceriesName := mealContent.MealGroceries
			quantityGram := mealContent.DishQuantityGram
			quantityUnit := mealContent.DishQuantityUnit

			groceries, exists := content.GM.GroceriesMap[groceriesName]
			if exists {
				pointerToGroceries := &groceries

				newMealContent = content.MealContents{
					MealGroceries:    pointerToGroceries,
					DishQuantityGram: quantityGram,
					DishQuantityUnit: quantityUnit,
				}

				println(quantityGram, quantityUnit)
			}
		} else if mealContent.MealDishes != "" {
			dishName := mealContent.MealDishes
			quantityGram := mealContent.DishQuantityGram
			quantityUnit := mealContent.DishQuantityUnit

			dish, exists := content.DM.DishMap[dishName]
			if exists {
				pointerToDish := &dish

				newMealContent = content.MealContents{
					MealDishes:       pointerToDish,
					DishQuantityGram: quantityGram,
					DishQuantityUnit: quantityUnit,
				}

			}
		}
	}

	var newMealData = content.Meal{
		Name:            mealData.Name,
		MealComposition: []content.MealContents{newMealContent},
	}

	err = content.MM.EditMealAttributeContent(newMealData.Name, newMealData.MealComposition)
	if err != nil {
		fmt.Println("Error adding meal:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Daten erfolgreich empfangen und verarbeitet!")

	println("Meal added!")

	for key, value := range content.MM.MealMap {
		fmt.Println(key, ":", value)
	}
}
