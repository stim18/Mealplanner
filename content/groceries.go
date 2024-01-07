// Matrikelnummern: 5911189 und 8441837
package content

import (
	"errors"
)

// Groceries with their attributes
type Groceries struct {
	Name            string
	Energy          float64
	Fat             float64
	SaturatedFats   float64
	Carbohydrates   float64
	Sugar           float64
	Protein         float64
	Salt            float64
	Fiber           float64
	Water           float64
	Weight          float64
	PackagingType   string
	PackagingWeight float64
}

// GroceriesManager adds all the groceries to a map
type GroceriesManager struct {
	GroceriesMap map[string]Groceries
}

var GM GroceriesManager

func init() {
	GM.GroceriesMap = make(map[string]Groceries)
}

// NewGroceries adds a Groceries to the map
func (gm *GroceriesManager) NewGroceries(groceries Groceries) error {
	// Check for empty input
	if groceries.Name == "" {
		return errors.New("Ungültiger Eingabewert: Lebensmittelname darf nicht leer sein")
	}

	if groceries.Weight == 0 && groceries.PackagingWeight == 0 {
		return errors.New("Fehlerhafte Eingabe: Entweder das Gewicht oder das Verpackungsgewicht mitgeben")
	}

	if groceries.Weight != 0 && groceries.PackagingWeight != 0 {
		return errors.New("Fehlerhafte Eingabe: Entweder das Gewicht oder das Verpackungsgewicht mitgeben")
	}

	// check, if Groceries already exists
	if _, exists := gm.GroceriesMap[groceries.Name]; exists {
		return errors.New("Lebensmittel mit diesem Namen existiert bereits")
	}

	// Check if Attributes are smaller than 0
	if groceries.Energy < 0 || groceries.Fat < 0 || groceries.SaturatedFats < 0 || groceries.Carbohydrates < 0 ||
		groceries.Sugar < 0 || groceries.Protein < 0 || groceries.Salt < 0 || groceries.Fiber < 0 ||
		groceries.Water < 0 || groceries.Weight < 0 || groceries.PackagingWeight < 0 {
		return errors.New("Fehlerhafte Eingabe: Die Attributwerte dürfen nicht kleiner als 0 sein")
	}

	// edit the input weight to 100 grams
	if groceries.Weight != 100.0 && groceries.PackagingWeight == 0 {
		factor := 100.0 / groceries.Weight

		groceries.Energy = groceries.Energy * factor
		groceries.Fat = groceries.Fat * factor
		groceries.SaturatedFats = groceries.SaturatedFats * factor
		groceries.Carbohydrates = groceries.Carbohydrates * factor
		groceries.Sugar = groceries.Sugar * factor
		groceries.Protein = groceries.Protein * factor
		groceries.Salt = groceries.Salt * factor
		groceries.Fiber = groceries.Fiber * factor
		groceries.Water = groceries.Water * factor
		groceries.Weight = groceries.Weight * factor
		groceries.PackagingType = ""
	}

	if groceries.PackagingWeight != 0 && groceries.PackagingType != "" && groceries.Weight == 0 {
		factor := 100 / groceries.PackagingWeight

		groceries.Energy = groceries.Energy * factor
		groceries.Fat = groceries.Fat * factor
		groceries.SaturatedFats = groceries.SaturatedFats * factor
		groceries.Carbohydrates = groceries.Carbohydrates * factor
		groceries.Sugar = groceries.Sugar * factor
		groceries.Protein = groceries.Protein * factor
		groceries.Salt = groceries.Salt * factor
		groceries.Fiber = groceries.Fiber * factor
		groceries.Water = groceries.Water * factor
		groceries.Weight = 100.0
	}

	// add the ingredient to the map
	gm.GroceriesMap[groceries.Name] = groceries
	return nil
}

// EditDishAttributes edits the attributes of a dish
func (gm *GroceriesManager) EditGroceriesName(name, newName string) error {
	// Check for empty input
	if name == "" || newName == "" {
		return errors.New("Ungültige Eingabewerte: Gerichtname darf nicht leer sein")
	}

	// Check if the dish exists
	groceries, exists := gm.GroceriesMap[name]
	if !exists {
		return errors.New("Gericht nicht gefunden")
	}

	// Check if there is another dish with the new Name
	if name != newName {
		if _, exists := gm.GroceriesMap[newName]; exists {
			return errors.New("Dish mit neuem Namen existiert bereits")
		}
	}

	// Update the Groceries Map
	delete(gm.GroceriesMap, name)
	groceries.Name = newName
	gm.GroceriesMap[newName] = groceries
	return nil
}

// EditGroceries updates a specific attribute of a food product
func (gm *GroceriesManager) EditGroceriesAttribute(groceries Groceries) error {
	// Check for empty input
	if groceries.Name == "" {
		return errors.New("Ungültiger Eingabewert: Lebensmittelname darf nicht leer sein")
	}

	// Checks if the given groceries exist
	_, exists := gm.GroceriesMap[groceries.Name]
	if !exists {
		return errors.New("Lebensmittel nicht gefunden")
	}

	err := GM.DeleteGroceries(groceries.Name)
	if err != nil {
		return err
	}

	err = GM.NewGroceries(groceries)
	if err != nil {
		return err
	}

	return nil
}

// DeleteGroceries removes a Groceries from the map
func (gm *GroceriesManager) DeleteGroceries(name string) error {
	// Check for empty input
	if name == "" {
		return errors.New("Ungültiger Eingabewert: Lebensmittelname darf nicht leer sein")
	}

	// Check whether the Groceries exist
	if _, exists := gm.GroceriesMap[name]; !exists {
		return errors.New("Lebensmittel nicht gefunden")
	}

	delete(gm.GroceriesMap, name)
	return nil
}
