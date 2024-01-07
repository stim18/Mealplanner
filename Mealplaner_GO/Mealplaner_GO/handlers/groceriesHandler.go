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
	"strconv"
)

func GroceriesHandler(w http.ResponseWriter, r *http.Request) {
	currentUser := data.GetCurrentUser()

	tmpl := template.Must(template.ParseFiles("./templates/groceries.html"))
	tmpl.Execute(w, currentUser)
}

func GetGroceries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(content.GM.GroceriesMap)
	if err != nil {
		http.Error(w, "Error encoding groceries", http.StatusInternalServerError)
		return
	}
}

func AddGroceriesHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Fehler beim Lesen des Anfragekörpers", http.StatusBadRequest)
		return
	}

	var groceryData map[string]interface{}
	err = json.Unmarshal([]byte(body), &groceryData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	energy := parseToFloat64(groceryData["Energy"])
	fat := parseToFloat64(groceryData["Fat"])
	saturatedFats := parseToFloat64(groceryData["SaturatedFats"])
	carbohydrates := parseToFloat64(groceryData["Carbohydrates"])
	sugar := parseToFloat64(groceryData["Sugar"])
	protein := parseToFloat64(groceryData["Protein"])
	salt := parseToFloat64(groceryData["Salt"])
	fiber := parseToFloat64(groceryData["Fiber"])
	water := parseToFloat64(groceryData["Water"])
	weight := parseToFloat64(groceryData["Weight"])
	packagingType := fmt.Sprintf("%v", groceryData["PackagingType"])
	packagingWeight := parseToFloat64(groceryData["PackagingWeight"])

	grocery := content.Groceries{
		Name:            groceryData["Name"].(string),
		Energy:          energy,
		Fat:             fat,
		SaturatedFats:   saturatedFats,
		Carbohydrates:   carbohydrates,
		Sugar:           sugar,
		Protein:         protein,
		Salt:            salt,
		Fiber:           fiber,
		Water:           water,
		Weight:          weight,
		PackagingType:   packagingType,
		PackagingWeight: packagingWeight,
	}

	err = content.GM.NewGroceries(grocery)
	if err != nil {
		fmt.Println("Error adding groceries:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Daten erfolgreich empfangen und verarbeitet!")
}

func EditGroceriesHandler(w http.ResponseWriter, r *http.Request) {

	println("moin")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Fehler beim Lesen des Anfragekörpers", http.StatusBadRequest)
		return
	}

	var groceryData map[string]interface{}
	err = json.Unmarshal([]byte(body), &groceryData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	energy := parseToFloat64(groceryData["Energy"])
	fat := parseToFloat64(groceryData["Fat"])
	saturatedFats := parseToFloat64(groceryData["SaturatedFats"])
	carbohydrates := parseToFloat64(groceryData["Carbohydrates"])
	sugar := parseToFloat64(groceryData["Sugar"])
	protein := parseToFloat64(groceryData["Protein"])
	salt := parseToFloat64(groceryData["Salt"])
	fiber := parseToFloat64(groceryData["Fiber"])
	water := parseToFloat64(groceryData["Water"])
	weight := parseToFloat64(groceryData["Weight"])
	packagingType := fmt.Sprintf("%v", groceryData["PackagingType"])
	packagingWeight := parseToFloat64(groceryData["PackagingWeight"])

	grocery := content.Groceries{
		Name:            groceryData["Name"].(string),
		Energy:          energy,
		Fat:             fat,
		SaturatedFats:   saturatedFats,
		Carbohydrates:   carbohydrates,
		Sugar:           sugar,
		Protein:         protein,
		Salt:            salt,
		Fiber:           fiber,
		Water:           water,
		Weight:          weight,
		PackagingType:   packagingType,
		PackagingWeight: packagingWeight,
	}

	err = content.GM.EditGroceriesAttribute(grocery)
	if err != nil {
		fmt.Println("Error editing groceries:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Daten erfolgreich empfangen und verarbeitet!")
}

func DeleteGroceriesHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Fehler beim Lesen des Anfragekörpers", http.StatusBadRequest)
		return
	}

	var groceryData map[string]interface{}
	err = json.Unmarshal([]byte(body), &groceryData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	groceryName := fmt.Sprintf("%v", groceryData["Name"])

	err = content.GM.DeleteGroceries(groceryName)
	if err != nil {
		fmt.Println("Error deleting groceries:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Daten erfolgreich empfangen und verarbeitet!")
}

func parseToFloat64(value interface{}) float64 {
	if value == "" {
		return 0
	}
	floatVal, err := strconv.ParseFloat(fmt.Sprintf("%v", value), 64)
	if err != nil {
		fmt.Println("Error parsing to float64:", err)
	}
	return floatVal
}
