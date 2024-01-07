// Matrikelnummern: 5911189 und 8441837
package content

import "errors"

// Day with content Date as identifier and Meals and NutritionalValues as Attributes
type Day struct {
	Date              string
	Meals             []Meal
	NutritionalValues Groceries
}

type DayManager struct {
	DayMap map[string]Day
}

var DAYM DayManager

func init() {
	DAYM.DayMap = make(map[string]Day)
}

// create a new Day
func (dm *DayManager) NewDay(date string, meals []Meal) error {
	if date == "" {
		return errors.New("Ungültiges Datum: Das Datum darf nicht leer sein")
	}
	if _, exists := dm.DayMap[date]; exists {
		return errors.New("Tag mit diesem Datum existiert bereits")
	}

	var newDay Day
	newDay.Date = date
	newDay.Meals = meals
	newDay.NutritionalValues = newDay.CalculateNutritionalValues()

	dm.DayMap[date] = newDay
	return nil
}

// delete existing Day
func (dm *DayManager) DeleteDay(date string) error {
	if _, exists := dm.DayMap[date]; !exists {
		return errors.New("Tag nicht gefunden")
	}

	delete(dm.DayMap, date)
	return nil
}

// edit an existing Day
func (dm *DayManager) EditDayDate(date string, newDate string) error {
	if date == "" || newDate == "" {
		return errors.New("Ungültige Eingabewerte: Das Datum darf nicht leer sein")
	}

	day, exists := dm.DayMap[date]
	if !exists {
		return errors.New("Tag nicht gefunden")
	}

	if date != newDate {
		if _, exists := dm.DayMap[newDate]; exists {
			return errors.New("Tag mit neuem Datum existiert bereits")
		}
	}

	delete(dm.DayMap, date)
	day.Date = newDate
	dm.DayMap[newDate] = day
	return nil
}

func (dm *DayManager) EditDayContents(date string, newMeals []Meal) error {
	if date == "" {
		return errors.New("Ungültige Eingabewerte: Datum darf nicht leer sein")
	}

	day, exists := dm.DayMap[date]
	if !exists {
		return errors.New("Tag nicht gefunden")
	}

	for _, meal := range newMeals {
		if meal.Name == "" || len(meal.MealComposition) == 0 {
			return errors.New("Ungültige Eingabe, Mahlzeit muss einen Namen und Inhalte haben")
		}
	}

	day.Meals = newMeals
	day.NutritionalValues = day.CalculateNutritionalValues()

	delete(dm.DayMap, date)
	dm.DayMap[date] = day
	return nil
}

func (d *Day) CalculateNutritionalValues() (nutritionalValues Groceries) {

	for _, meal := range d.Meals {
		nutritionalValues.Energy += meal.NutritionalValues.Energy
		nutritionalValues.Fat += meal.NutritionalValues.Fat
		nutritionalValues.SaturatedFats += meal.NutritionalValues.SaturatedFats
		nutritionalValues.Carbohydrates += meal.NutritionalValues.Carbohydrates
		nutritionalValues.Sugar += meal.NutritionalValues.Sugar
		nutritionalValues.Protein += meal.NutritionalValues.Protein
		nutritionalValues.Salt += meal.NutritionalValues.Salt
		nutritionalValues.Fiber += meal.NutritionalValues.Fiber
		nutritionalValues.Water += meal.NutritionalValues.Water
	}
	return nutritionalValues
}
