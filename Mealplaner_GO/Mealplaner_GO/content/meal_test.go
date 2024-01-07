// Matrikelnummern: 5911189 und 8441837
package content

import (
	"reflect"
	"testing"
)

func TestNewMeal(t *testing.T) {
	mealMgr := MealManager{
		MealMap: make(map[string]Meal),
	}

	// Testfall: Hinzufügen einer neuen Mahlzeit mit einem Lebensmittel und einem Gericht
	t.Run("AddNewMealWithGroceryAndDish", func(t *testing.T) {
		// Neues Gericht und Lebensmittel hinzufügen
		banana := Groceries{Name: "Banana", Energy: 89, Fat: 0.3, Carbohydrates: 23, Protein: 1.1}
		yogurt := Groceries{Name: "Yogurt", Energy: 61, Fat: 3.3, Carbohydrates: 4.7, Protein: 3.5}

		// Erstellen des DishContents für das Gericht
		bananaContents := DishContents{
			DishGroceries:    &banana,
			DishDishes:       nil,
			DishQuantityGram: 100,
			DishQuantityUnit: 0,
		}

		yogurtContents := DishContents{
			DishGroceries:    &yogurt,
			DishDishes:       nil,
			DishQuantityGram: 150,
			DishQuantityUnit: 0,
		}
		bananaDishComposition := []DishContents{bananaContents, yogurtContents}

		ingredients := []MealContents{
			{
				MealGroceries:    &Groceries{Name: "Apple"},
				MealDishes:       nil,
				DishQuantityGram: 200.0,
				DishQuantityUnit: 0,
			},
			{
				MealGroceries:    nil,
				MealDishes:       &Dish{Name: "Banana Dish", DishComposition: bananaDishComposition},
				DishQuantityGram: 0,
				DishQuantityUnit: 1,
			},
		}

		err := mealMgr.NewMeal("AppleBananaMeal", ingredients)
		if err != nil {
			t.Errorf("Expected no error, got %s", err.Error())
		}

		// Überprüfen, ob die Mahlzeit zur Mahlzeiten-Map hinzugefügt wurde
		_, exists := mealMgr.MealMap["AppleBananaMeal"]
		if !exists {
			t.Errorf("Expected meal to be added, but it's not present in the map")
		}
	})

	// Testfall 2: Hinzufügen einer Mahlzeit mit leerem Namen
	t.Run("AddMealWithEmptyName", func(t *testing.T) {
		// Leeres Gericht hinzufügen
		err := mealMgr.NewMeal("", nil)
		expectedErrMsg := "Ungültiger Eingabewert: Mahlzeitname darf nicht leer sein"

		if err == nil || err.Error() != expectedErrMsg {
			t.Errorf("Expected error: %s, got %v", expectedErrMsg, err)
		}
	})

	// Testfall 3: Hinzufügen einer bereits vorhandenen Mahlzeit
	t.Run("AddDuplicateMeal", func(t *testing.T) {
		// Vorhandene Mahlzeit hinzufügen
		ingredients := []MealContents{
			{
				MealGroceries:    &Groceries{Name: "Banana"},
				MealDishes:       nil,
				DishQuantityGram: 150.0,
				DishQuantityUnit: 0,
			},
		}

		mealMgr.NewMeal("AppleMeal", ingredients)
		err := mealMgr.NewMeal("AppleMeal", ingredients)
		expectedErrMsg := "Mahlzeit mit diesem Namen existiert bereits"

		if err == nil || err.Error() != expectedErrMsg {
			t.Errorf("Expected error: %s, got %v", expectedErrMsg, err)
		}
	})
}

func TestDeleteMeal(t *testing.T) {
	mealMgr := MealManager{
		MealMap: map[string]Meal{
			"Meal1": {Name: "Meal1"},
			"Meal2": {Name: "Meal2"},
		},
	}

	// Testfall 1: Löschen einer vorhandenen Mahlzeit
	t.Run("DeleteExistingMeal", func(t *testing.T) {
		err := mealMgr.DeleteMeal("Meal1")
		if err != nil {
			t.Errorf("Expected no error, got %s", err.Error())
		}

		// Überprüfen, ob die Mahlzeit gelöscht wurde
		_, exists := mealMgr.MealMap["Meal1"]
		if exists {
			t.Errorf("Expected meal to be deleted, but it's still present in the map")
		}
	})

	// Testfall 2: Löschen einer nicht vorhandenen Mahlzeit
	t.Run("DeleteNonExistingMeal", func(t *testing.T) {
		err := mealMgr.DeleteMeal("NonExistingMeal")
		expectedErrMsg := "Mahlzeit nicht gefunden"

		if err == nil || err.Error() != expectedErrMsg {
			t.Errorf("Expected error: %s, got %v", expectedErrMsg, err)
		}
	})
}

func TestEditMealAttributes(t *testing.T) {
	groceriesTest := Groceries{
		Name:            "Gurkin",
		Energy:          52.0,
		Fat:             0.2,
		SaturatedFats:   0.02,
		Carbohydrates:   13.8,
		Sugar:           10.4,
		Protein:         0.3,
		Salt:            0.0,
		Fiber:           2.4,
		Water:           85.6,
		Weight:          100.0,
		PackagingWeight: 5.0,
	}

	mm := &MealManager{
		MealMap: map[string]Meal{
			"Dinner": {
				Name: "Dinner",
				MealComposition: []MealContents{
					{
						MealGroceries:    &groceriesTest,
						MealDishes:       nil,
						DishQuantityGram: 200.0,
						DishQuantityUnit: 0,
					},
				},
				// values added manually because of rounding issues
				NutritionalValues: Groceries{
					Energy:        156,
					Fat:           0.6000000000000001,
					SaturatedFats: 0.06,
					Carbohydrates: 41.400000000000006,
					Sugar:         31.200000000000003,
					Protein:       0.8999999999999999,
					Salt:          0,
					Fiber:         7.199999999999999,
					Water:         256.79999999999995,
				},
			},
		},
	}

	tests := []struct {
		name        string
		newName     string
		newContents []MealContents
		expectedErr error
		expectedMap map[string]Meal
	}{
		{
			name:        "Dinner",
			newName:     "Supper",
			newContents: []MealContents{{MealGroceries: &groceriesTest, DishQuantityGram: 300.0, DishQuantityUnit: 0}},
			expectedErr: nil,
			expectedMap: map[string]Meal{
				"Supper": {
					Name: "Supper",
					MealComposition: []MealContents{
						{
							MealGroceries:    &groceriesTest,
							DishQuantityGram: 300.0,
							DishQuantityUnit: 0,
						},
					},
					NutritionalValues: Groceries{
						Energy:        156,
						Fat:           0.6000000000000001,
						SaturatedFats: 0.06,
						Carbohydrates: 41.400000000000006,
						Sugar:         31.200000000000003,
						Protein:       0.8999999999999999,
						Salt:          0,
						Fiber:         7.199999999999999,
						Water:         256.79999999999995,
					},
				},
			},
		},
	}
	// Test if the content and name edited successfully
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := mm.EditMealName(tt.name, tt.newName)
			err = mm.EditMealAttributeContent(tt.newName, tt.newContents)
			if !reflect.DeepEqual(err, tt.expectedErr) {
				t.Errorf("Got error %v, but expected %v", err, tt.expectedErr)
			}

			if !reflect.DeepEqual(mm.MealMap, tt.expectedMap) {
				t.Errorf("Got map %v, but expected %v", mm.MealMap, tt.expectedMap)
			}
		})
	}
}
