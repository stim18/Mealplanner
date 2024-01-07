// Matrikelnummern: 5911189 und 8441837
package userManagement

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"meinProjekt/content"
	"meinProjekt/data"
)

func LoadUserData(username string) error {
	filepath := fmt.Sprintf("data/%s.json", username)

	// JSON-Datei einlesen
	encryptedData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil
	}

	// Entschlüsselung der Daten
	decryptedData, err := Decrypt(encryptedData)
	if err != nil {
		return err
	}

	// Daten aus JSON in ein Interface-Objekt umwandeln
	var data map[string]interface{}
	if err := json.Unmarshal(decryptedData, &data); err != nil {
		return err
	}

	for _, value := range data["groceriesMap"].(map[string]interface{}) {
		grocery := convertToGroceries(value)
		err := content.GM.NewGroceries(grocery)
		if err != nil {
			fmt.Println("Fehler beim Hinzufügen eines neuen Lebensmittels:", err)
		}
	}

	for _, value := range data["dishMap"].(map[string]interface{}) {
		dish := convertToDish(value)
		err := content.DM.NewDish(dish.Name, dish.DishComposition)
		if err != nil {
			fmt.Println("Fehler beim Hinzufügen eines neuen Gerichts:", err)
		}
	}

	for _, value := range data["mealMap"].(map[string]interface{}) {
		meal := convertToMeal(value)
		err := content.MM.NewMeal(meal.Name, meal.MealComposition)
		if err != nil {
			fmt.Println("Fehler beim Hinzufügen einer neuen Mahlzeit:", err)
		}
	}

	for _, value := range data["dayMap"].(map[string]interface{}) {
		day := convertToDay(value)
		err := content.DAYM.NewDay(day.Date, day.Meals)
		if err != nil {
			fmt.Println("Fehler beim Hinzufügen eines neuen Tags:", err)
		}
	}

	return nil
}

func convertToDay(data interface{}) (day content.Day) {
	dayData := data.(map[string]interface{})

	date := dayData["Date"].(string)
	mealsRaw := dayData["Meals"].([]interface{})

	var meals []content.Meal

	for _, rawMeal := range mealsRaw {
		convertedMeal := convertToMeal(rawMeal)
		meals = append(meals, convertedMeal)
	}

	nutritionalValues := convertToGroceries(dayData["NutritionalValues"].(map[string]interface{}))

	day = content.Day{
		Date:              date,
		Meals:             meals,
		NutritionalValues: nutritionalValues,
	}

	return day
}

func Decrypt(ciphertext []byte) ([]byte, error) {
	currentUser := data.GetCurrentUser()
	passphrase := []byte(currentUser.Password)

	block, err := aes.NewCipher(passphrase)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext))
	stream.XORKeyStream(decrypted, ciphertext)
	return decrypted, nil
}

func convertToMeal(data interface{}) (meal content.Meal) {
	mealData := data.(map[string]interface{})

	name := mealData["Name"].(string)
	mealCompositionRaw := mealData["MealComposition"].([]interface{})

	var mealComposition []content.MealContents

	for _, rawContent := range mealCompositionRaw {
		convertedContent := rawContent.(map[string]interface{})

		mealContent := content.MealContents{
			DishQuantityGram: convertedContent["DishQuantityGram"].(float64),
			DishQuantityUnit: convertedContent["DishQuantityUnit"].(float64),
		}

		// Setze MealGroceries, falls vorhanden
		if rawGroceries, ok := convertedContent["MealGroceries"]; ok && rawGroceries != nil {
			groceries := convertToGroceries(rawGroceries)
			mealContent.MealGroceries = &groceries
		}

		// Setze MealDishes, falls vorhanden
		if rawDish, ok := convertedContent["MealDishes"]; ok && rawDish != nil {
			dish := convertToDish(rawDish)
			mealContent.MealDishes = &dish
		}

		mealComposition = append(mealComposition, mealContent)
	}

	nutritionalValues := convertToGroceries(mealData["NutritionalValues"].(map[string]interface{}))

	meal = content.Meal{
		Name:              name,
		MealComposition:   mealComposition,
		NutritionalValues: nutritionalValues,
	}

	return meal
}

func convertToDish(data interface{}) (dish content.Dish) {
	dishData := data.(map[string]interface{})

	name := dishData["Name"].(string)
	// Rohdaten der DishComposition
	dishCompositionRaw := dishData["DishComposition"].([]interface{})

	var dishComposition []content.DishContents

	// Iteriere über die Rohdaten der DishComposition
	for _, rawContent := range dishCompositionRaw {
		convertedContent := rawContent.(map[string]interface{})
		dishContent := content.DishContents{
			DishQuantityGram: convertedContent["DishQuantityGram"].(float64),
			DishQuantityUnit: convertedContent["DishQuantityUnit"].(float64),
		}
		if rawGroceries, ok := convertedContent["DishGroceries"]; ok && rawGroceries != nil {
			groceries := convertToGroceries(rawGroceries)
			dishContent.DishGroceries = &groceries
		}
		if rawDish, ok := convertedContent["DishDishes"]; ok && rawDish != nil {
			dish := convertToDish(rawDish)
			dishContent.DishDishes = &dish
		}
		dishComposition = append(dishComposition, dishContent)
	}

	nutritionalValues := convertToGroceries(dishData["NutritionalValues"].(map[string]interface{}))

	dish = content.Dish{
		Name:              name,
		DishComposition:   dishComposition,
		NutritionalValues: nutritionalValues,
	}

	return dish
}

func convertToGroceries(data interface{}) (groceries content.Groceries) {

	groceriesData := data.(map[string]interface{})

	name := groceriesData["Name"].(string)
	energy := groceriesData["Energy"].(float64)
	fat := groceriesData["Fat"].(float64)
	saturatedFats := groceriesData["SaturatedFats"].(float64)
	carbohydrates := groceriesData["Carbohydrates"].(float64)
	sugar := groceriesData["Sugar"].(float64)
	protein := groceriesData["Protein"].(float64)
	salt := groceriesData["Salt"].(float64)
	fiber := groceriesData["Fiber"].(float64)
	water := groceriesData["Water"].(float64)
	weight := groceriesData["Weight"].(float64)
	packagingType := groceriesData["PackagingType"].(string)
	packagingWeight := groceriesData["PackagingWeight"].(float64)

	groceries = content.Groceries{
		Name:            name,
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
	return groceries
}
