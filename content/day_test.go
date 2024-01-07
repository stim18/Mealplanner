// Matrikelnummern: 5911189 und 8441837
package content

import (
	"reflect"
	"testing"
)

func TestNewDay(t *testing.T) {
	dm := &DayManager{
		DayMap: make(map[string]Day),
	}

	meals := []Meal{
		{
			Name: "Breakfast",
		},
		{
			Name: "Lunch",
		},
	}

	// Test Case 1: Add new Day
	err := dm.NewDay("2024-01-03", meals)
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	// Check if new Day was added to map
	_, exists := dm.DayMap["2024-01-03"]
	if !exists {
		t.Errorf("Expected day to be added, but it's not present in the map")
	}

	// Test Case 2: Add an already existing Day
	err = dm.NewDay("2024-01-03", meals)
	expectedErrMsg := "Tag mit diesem Datum existiert bereits"
	if err == nil || err.Error() != expectedErrMsg {
		t.Errorf("Expected error: %s, got %v", expectedErrMsg, err)
	}

	// Test Case 3: Add a Day without a date
	err = dm.NewDay("", meals)
	expectedErrMsg = "Ungültiges Datum: Das Datum darf nicht leer sein"
	if err == nil || err.Error() != expectedErrMsg {
		t.Errorf("Expected error: %s, got %v", expectedErrMsg, err)
	}
}

func TestDeleteDay(t *testing.T) {
	dm := &DayManager{
		DayMap: make(map[string]Day),
	}

	meals := []Meal{
		{
			Name: "Breakfast",
		},
		{
			Name: "Lunch",
		},
	}

	// Add new day for tests
	err := dm.NewDay("2024-01-03", meals)
	if err != nil {
		t.Errorf("Error while adding day: %s", err.Error())
	}

	// Test Case: delete day
	err = dm.DeleteDay("2024-01-03")
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	// Check if date got deleted out of map
	_, exists := dm.DayMap["2024-01-03"]
	if exists {
		t.Errorf("Expected day to be deleted, but it's still present in the map")
	}

	// Test Case: delete a day that does not exist
	err = dm.DeleteDay("2024-01-04")
	expectedErrMsg := "Tag nicht gefunden"
	if err == nil || err.Error() != expectedErrMsg {
		t.Errorf("Expected error: %s, got %v", expectedErrMsg, err)
	}
}

func TestEditMealAttributeContent(t *testing.T) {
	mm := &MealManager{
		MealMap: make(map[string]Meal),
	}

	// Add a Meal for Test
	contents := []MealContents{
		{
			MealGroceries:    &Groceries{Name: "Apple", Energy: 52.0, Protein: 0.3},
			MealDishes:       nil,
			DishQuantityGram: 200.0,
			DishQuantityUnit: 0,
		},
	}
	err := mm.NewMeal("AppleMeal", contents)
	if err != nil {
		t.Errorf("Error setting up the test: %s", err.Error())
	}

	newContents := []MealContents{
		{
			MealGroceries:    &Groceries{Name: "Banana", Energy: 89.0, Protein: 1.1},
			MealDishes:       nil,
			DishQuantityGram: 150.0,
			DishQuantityUnit: 0,
		},
	}

	err = mm.EditMealAttributeContent("AppleMeal", newContents)
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	updatedMeal, exists := mm.MealMap["AppleMeal"]
	if !exists {
		t.Errorf("Expected meal to be updated, but it's not present in the map")
	}

	if len(updatedMeal.MealComposition) != 1 {
		t.Errorf("Expected meal composition length to be 1, got %d", len(updatedMeal.MealComposition))
	}

	expectedGrocery := &Groceries{Name: "Banana", Energy: 89.0, Protein: 1.1}
	if !reflect.DeepEqual(*updatedMeal.MealComposition[0].MealGroceries, *expectedGrocery) {
		t.Errorf("Expected meal content to be updated, got %v", updatedMeal.MealComposition[0].MealGroceries)
	}
}

func TestEditDayContents(t *testing.T) {
	dm := &DayManager{
		DayMap: make(map[string]Day),
	}

	// Add Meal to Day
	meals := []Meal{
		{
			Name: "Breakfast",
			MealComposition: []MealContents{
				{
					MealGroceries:    &Groceries{Name: "Apple", Energy: 52.0, Protein: 0.3},
					MealDishes:       nil,
					DishQuantityGram: 200.0,
					DishQuantityUnit: 0,
				},
			},
		},
	}
	err := dm.NewDay("2024-01-03", meals)
	if err != nil {
		t.Errorf("Error setting up the test: %s", err.Error())
	}

	newMeals := []Meal{
		{
			Name: "Lunch",
			MealComposition: []MealContents{
				{
					MealGroceries:    &Groceries{Name: "Banana", Energy: 89.0, Protein: 1.1},
					MealDishes:       nil,
					DishQuantityGram: 150.0,
					DishQuantityUnit: 0,
				},
			},
		},
	}

	err = dm.EditDayContents("2024-01-03", newMeals)
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	updatedDay, exists := dm.DayMap["2024-01-03"]
	if !exists {
		t.Errorf("Expected day to be updated, but it's not present in the map")
	}

	if len(updatedDay.Meals) != 1 {
		t.Errorf("Expected number of meals to be 1, got %d", len(updatedDay.Meals))
	}

	expectedGrocery := &Groceries{Name: "Banana", Energy: 89.0, Protein: 1.1}
	if !reflect.DeepEqual(*updatedDay.Meals[0].MealComposition[0].MealGroceries, *expectedGrocery) {
		t.Errorf("Expected meal content to be updated, got %v", updatedDay.Meals[0].MealComposition[0].MealGroceries)
	}
}

func TestCalculateNutritionalValuesDay(t *testing.T) {
	meal1 := Meal{
		Name: "Breakfast",
		NutritionalValues: Groceries{
			Energy:        200.0,
			Fat:           10.0,
			SaturatedFats: 3.0,
			Carbohydrates: 30.0,
			Sugar:         15.0,
			Protein:       25.0,
			Salt:          2.0,
			Fiber:         5.0,
			Water:         100.0,
		},
	}
	meal2 := Meal{
		Name: "Lunch",
		NutritionalValues: Groceries{
			Energy:        300.0,
			Fat:           15.0,
			SaturatedFats: 4.0,
			Carbohydrates: 40.0,
			Sugar:         20.0,
			Protein:       30.0,
			Salt:          2.5,
			Fiber:         6.0,
			Water:         150.0,
		},
	}

	day := Day{
		Date:  "2024-01-03",
		Meals: []Meal{meal1, meal2},
	}

	nutritionalValues := day.CalculateNutritionalValues()

	expectedValues := Groceries{
		Energy:        500.0,
		Fat:           25.0,
		SaturatedFats: 7.0,
		Carbohydrates: 70.0,
		Sugar:         35.0,
		Protein:       55.0,
		Salt:          4.5,
		Fiber:         11.0,
		Water:         250.0,
	}

	if !reflect.DeepEqual(nutritionalValues, expectedValues) {
		t.Errorf("Expected nutritional values to be %v, got %v", expectedValues, nutritionalValues)
	}
}
