// Matrikelnummern: 5911189 und 8441837
package content

import (
	"errors"
)

// Meal with content Name as identifier and MealComposition and NutritionalValues as Attributes
type Meal struct {
	Name              string
	MealComposition   []MealContents
	NutritionalValues Groceries
}

// Contents of Meal can be MealGroceries, MealDishes, DishQuantityGram and DishQuantityUnit
type MealContents struct {
	MealGroceries    *Groceries
	MealDishes       *Dish
	DishQuantityGram float64
	DishQuantityUnit float64
}

type MealManager struct {
	MealMap map[string]Meal
}

var MM MealManager

func init() {
	MM.MealMap = make(map[string]Meal)
}

// create a New Meal with an input Name and Contents of MealContents
func (mm *MealManager) NewMeal(nameMeal string, newContents []MealContents) error {
	if nameMeal == "" {
		return errors.New("Ungültiger Eingabewert: Mahlzeitname darf nicht leer sein")
	}
	if _, exists := mm.MealMap[nameMeal]; exists {
		return errors.New("Mahlzeit mit diesem Namen existiert bereits")
	}
	for _, content := range newContents {
		if content.MealGroceries == nil && content.MealDishes == nil {
			return errors.New("Ungültige Eingabe, weder Lebensmittel noch Mahlzeit angegeben")
		}
		if content.DishQuantityGram == 0.0 && content.DishQuantityUnit == 0.0 {
			return errors.New("Ungültige Eingabe, weder Gramm noch Anzahl angegeben")
		}
		if content.MealGroceries != nil && content.MealDishes != nil {
			return errors.New("Ungültige Eingabe, Lebensmittel oder Mahlzeit angeben")
		}
		if content.DishQuantityGram != 0.0 && content.DishQuantityUnit != 0.0 {
			return errors.New("Ungültige Eingabe, Gramm oder Anzahl angeben")
		}
	}

	var newMeal Meal
	newMeal.Name = nameMeal
	newMeal.MealComposition = newContents
	newMeal.NutritionalValues = newMeal.CalculateNutritionalValues()

	mm.MealMap[nameMeal] = newMeal
	return nil
}

// delete existing Meal
func (mm *MealManager) DeleteMeal(name string) error {
	if _, exists := mm.MealMap[name]; !exists {
		return errors.New("Mahlzeit nicht gefunden")
	}

	delete(mm.MealMap, name)
	return nil
}

// edit the name of a Meal
func (mm *MealManager) EditMealName(name string, newName string) error {
	if name == "" || newName == "" {
		return errors.New("Ungültige Eingabewerte: Mahlzeitname darf nicht leer sein")
	}

	meal, exists := mm.MealMap[name]
	if !exists {
		return errors.New("Mahlzeit nicht gefunden")
	}

	if name != newName {
		if _, exists := mm.MealMap[newName]; exists {
			return errors.New("Mahlzeit mit neuem Namen existiert bereits")
		}
	}

	delete(mm.MealMap, name)
	meal.Name = newName
	mm.MealMap[newName] = meal
	return nil
}

// Edit the content of a meal
func (mm *MealManager) EditMealAttributeContent(name string, newContents []MealContents) error {
	if name == "" {
		return errors.New("Ungültige Eingabewerte: Mahlzeitname darf nicht leer sein")
	}

	meal, exists := mm.MealMap[name]
	if !exists {
		return errors.New("Mahlzeit nicht gefunden")
	}

	for _, content := range newContents {
		if content.MealGroceries == nil && content.MealDishes == nil {
			return errors.New("Ungültige Eingabe, weder Lebensmittel noch Mahlzeit angegeben")
		}
		if content.DishQuantityGram == 0.0 && content.DishQuantityUnit == 0.0 {
			return errors.New("Ungültige Eingabe, weder Gramm noch Anzahl angegeben")
		}
		if content.MealGroceries != nil && content.MealDishes != nil {
			return errors.New("Ungültige Eingabe, Lebensmittel oder Mahlzeit angeben")
		}
		if content.DishQuantityGram != 0.0 && content.DishQuantityUnit != 0.0 {
			return errors.New("Ungültige Eingabe, Gramm oder Anzahl angeben")
		}
	}

	meal.MealComposition = newContents
	// calculate the nutritionalValues of a meal before saving it and safe them as an attribute of the meal

	mm.DeleteMeal(name)
	mm.NewMeal(name, meal.MealComposition)
	return nil
}

// calculate the nutritions of a meal
func (m *Meal) CalculateNutritionalValues() (nutritionalValues Groceries) {
	for _, ingredients := range m.MealComposition {
		if ingredients.MealGroceries != nil {

			nutritionalValues.Energy += ingredients.MealGroceries.Energy * (ingredients.DishQuantityGram / 100.0)
			nutritionalValues.Fat += ingredients.MealGroceries.Fat * (ingredients.DishQuantityGram / 100.0)
			nutritionalValues.SaturatedFats += ingredients.MealGroceries.SaturatedFats * (ingredients.DishQuantityGram / 100.0)
			nutritionalValues.Carbohydrates += ingredients.MealGroceries.Carbohydrates * (ingredients.DishQuantityGram / 100.0)
			nutritionalValues.Sugar += ingredients.MealGroceries.Sugar * (ingredients.DishQuantityGram / 100.0)
			nutritionalValues.Protein += ingredients.MealGroceries.Protein * (ingredients.DishQuantityGram / 100.0)
			nutritionalValues.Salt += ingredients.MealGroceries.Salt * (ingredients.DishQuantityGram / 100.0)
			nutritionalValues.Fiber += ingredients.MealGroceries.Fiber * (ingredients.DishQuantityGram / 100.0)
			nutritionalValues.Water += ingredients.MealGroceries.Water * (ingredients.DishQuantityGram / 100.0)

		} else if ingredients.MealDishes != nil {

			dishNutritionalValues := ingredients.MealDishes.CalculateNutritionalValues()

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
