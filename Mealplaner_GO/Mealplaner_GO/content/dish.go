// Matrikelnummern: 5911189 und 8441837
package content

import (
	"errors"
)

// Dish with Attributes Name, DishComposition, NutritionalValues
type Dish struct {
	Name              string
	DishComposition   []DishContents
	NutritionalValues Groceries
}

// DishContents reflects the content of a  Dish
type DishContents struct {
	DishGroceries    *Groceries
	DishDishes       *Dish
	DishQuantityGram float64
	DishQuantityUnit float64
}

// DishManager creates a map to save Dishes
type DishManager struct {
	DishMap map[string]Dish
}

var DM DishManager

func init() {
	DM.DishMap = make(map[string]Dish)
}

// create a new dish
func (dm *DishManager) NewDish(nameDish string, newContents []DishContents) error {
	// check for invalid input
	if nameDish == "" {
		return errors.New("Ungültiger Eingabewert: Gerichtname darf nicht leer sein")
	}
	//Check if the name already exists
	if _, exists := dm.DishMap[nameDish]; exists {
		return errors.New("Gericht mit diesem Namen existiert bereits")
	}
	for _, content := range newContents {
		if content.DishGroceries == nil && content.DishDishes == nil {
			return errors.New("Ungültige Eingabe, weder Lebensmittel noch Mahlzeit angegeben")
		}
		if content.DishQuantityGram == 0.0 && content.DishQuantityUnit == 0.0 {
			return errors.New("Ungültige Eingabe, weder Gramm noch Anzahl angegeben")
		}
		if content.DishGroceries != nil && content.DishDishes != nil {
			return errors.New("Ungültige Eingabe, Lebensmittel oder Mahlzeit angeben")
		}
		if content.DishQuantityGram != 0.0 && content.DishQuantityUnit != 0.0 {
			return errors.New("Ungültige Eingabe, Gramm oder Anzahl angeben")
		}
	}
	var newDish Dish
	newDish.Name = nameDish
	newDish.DishComposition = newContents
	newDish.NutritionalValues = newDish.CalculateNutritionalValues()

	//Adding the dish to the map
	dm.DishMap[nameDish] = newDish
	return nil
}

// EditDishAttributes edits the attributes of a dish
func (dm *DishManager) EditDishAttributeName(name, newName string) error {
	// Check for empty input
	if name == "" || newName == "" {
		return errors.New("Ungültige Eingabewerte: Gerichtname darf nicht leer sein")
	}

	// Check if the dish exists
	dish, exists := dm.DishMap[name]
	if !exists {
		return errors.New("Gericht nicht gefunden")
	}

	// Check if there is another dish with the new Name
	if name != newName {
		if _, exists := dm.DishMap[newName]; exists {
			return errors.New("Dish mit neuem Namen existiert bereits")
		}
	}

	// refresh the Map with new input
	delete(dm.DishMap, name)
	dish.Name = newName
	dish.NutritionalValues = dish.CalculateNutritionalValues()
	dm.DishMap[newName] = dish
	return nil
}

// EditDishAttributes edits the attributes of a dish
func (dm *DishManager) EditDishAttributeContent(name string, newContents []DishContents) error {
	// Check for empty input
	if name == "" {
		return errors.New("Ungültige Eingabewerte: Gerichtname darf nicht leer sein")
	}

	// Check if the dish exists
	dish, exists := dm.DishMap[name]
	if !exists {
		return errors.New("Gericht nicht gefunden")
	}

	for _, content := range newContents {
		if content.DishGroceries == nil && content.DishDishes == nil {
			return errors.New("Ungültiger Typ für DishContents")
		}
		if content.DishQuantityGram == 0.0 && content.DishQuantityUnit == 0.0 {
			return errors.New("Ungültiger Typ für DishContents")
		}
		if content.DishGroceries != nil && content.DishDishes != nil {
			return errors.New("Ungültiger Typ für DishContents")
		}
		if content.DishQuantityGram != 0.0 && content.DishQuantityUnit != 0.0 {
			return errors.New("Ungültiger Typ für DishContents")
		}
	}

	// Update dish attributes
	dish.DishComposition = newContents
	dish.NutritionalValues = dish.CalculateNutritionalValues()

	// Update Dish
	delete(dm.DishMap, name)
	dm.DishMap[name] = dish
	return nil
}

// Removing a dish from the map
func (dm *DishManager) DeleteDish(name string) error {
	//Check if the dish exists
	if _, exists := dm.DishMap[name]; !exists {
		return errors.New("Gericht nicht gefunden")
	}

	delete(dm.DishMap, name)
	return nil
}

// computes the cumulative nutritional values of a Dish
func (d *Dish) CalculateNutritionalValues() (nutritionalValues Groceries) {
	// Iterate through each ingredient in the Dish's composition
	for _, ingredients := range d.DishComposition {
		if ingredients.DishGroceries != nil {
			// Calculate nutritional values based on quantity in grams and accumulate them into the overall nutritionalValues
			nutritionalValues.Energy += ingredients.DishGroceries.Energy * (ingredients.DishQuantityGram / 100.0)
			nutritionalValues.Fat += ingredients.DishGroceries.Fat * (ingredients.DishQuantityGram / 100.0)
			nutritionalValues.SaturatedFats += ingredients.DishGroceries.SaturatedFats * (ingredients.DishQuantityGram / 100.0)
			nutritionalValues.Carbohydrates += ingredients.DishGroceries.Carbohydrates * (ingredients.DishQuantityGram / 100.0)
			nutritionalValues.Sugar += ingredients.DishGroceries.Sugar * (ingredients.DishQuantityGram / 100.0)
			nutritionalValues.Protein += ingredients.DishGroceries.Protein * (ingredients.DishQuantityGram / 100.0)
			nutritionalValues.Salt += ingredients.DishGroceries.Salt * (ingredients.DishQuantityGram / 100.0)
			nutritionalValues.Fiber += ingredients.DishGroceries.Fiber * (ingredients.DishQuantityGram / 100.0)
			nutritionalValues.Water += ingredients.DishGroceries.Water * (ingredients.DishQuantityGram / 100.0)
		} else if ingredients.DishDishes != nil {
			// Calculate nutritional values based on quantity of units accumulate them into the overall nutritionalValues
			dishNutritionalValues := ingredients.DishDishes.CalculateNutritionalValues()

			nutritionalValues.Energy += dishNutritionalValues.Energy * (ingredients.DishQuantityUnit)
			nutritionalValues.Fat += dishNutritionalValues.Fat * (ingredients.DishQuantityUnit)
			nutritionalValues.SaturatedFats += dishNutritionalValues.SaturatedFats * (ingredients.DishQuantityUnit)
			nutritionalValues.Carbohydrates += dishNutritionalValues.Carbohydrates * (ingredients.DishQuantityUnit)
			nutritionalValues.Sugar += dishNutritionalValues.Sugar * (ingredients.DishQuantityUnit)
			nutritionalValues.Protein += dishNutritionalValues.Protein * (ingredients.DishQuantityUnit)
			nutritionalValues.Salt += dishNutritionalValues.Salt * (ingredients.DishQuantityUnit)
			nutritionalValues.Fiber += dishNutritionalValues.Fiber * (ingredients.DishQuantityUnit)
			nutritionalValues.Water += dishNutritionalValues.Water * (ingredients.DishQuantityUnit)
		}
	}
	return nutritionalValues
}
