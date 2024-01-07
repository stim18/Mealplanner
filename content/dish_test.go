// Matrikelnummern: 5911189 und 8441837
package content

import (
	"reflect"
	"testing"
)

func TestNewDish(t *testing.T) {
	dm := &DishManager{
		DishMap: make(map[string]Dish),
	}

	// Dish normal Case
	t.Run("NewDish_NormalCase", func(t *testing.T) {
		groceries := Groceries{
			Name:            "Apple",
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
		dish := Dish{
			Name: "Apple Dish",
			DishComposition: []DishContents{
				{
					DishGroceries:    &groceries,
					DishQuantityGram: 150.0,
				},
			},
		}

		groceries1 := Groceries{
			Name:            "Banana",
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
		dish1 := Dish{
			Name: "Banana Dish",
			DishComposition: []DishContents{
				{
					DishGroceries:    &groceries1,
					DishQuantityGram: 150.0,
				},
			},
		}

		err := dm.NewDish(dish.Name, dish.DishComposition)
		err = dm.NewDish(dish1.Name, dish.DishComposition)
		if err != nil {
			t.Errorf("Expected no error, got ")
		}
	})

	t.Run("NewDish_MissingAttributes", func(t *testing.T) {

		// DishComposition is empty
		dish := Dish{
			Name: "Strawberry Dish",
			DishComposition: []DishContents{{
				DishGroceries:    nil,
				DishDishes:       nil,
				DishQuantityGram: 200.0,
				DishQuantityUnit: 0,
			},
			},
		}

		err := dm.NewDish(dish.Name, dish.DishComposition)
		if err == nil {
			t.Errorf("Error expected, got nil")
		} else {
			expectedErrMsgOne := "Ungültige Eingabe, weder Lebensmittel noch Mahlzeit angegeben"
			expectedErrMsgTwo := "Ungültige Eingabe, weder Gramm noch Anzahl angegeben"
			if err.Error() != expectedErrMsgOne && err.Error() != expectedErrMsgTwo {
				t.Errorf("Expected error: %s or %s, got %v", expectedErrMsgOne, expectedErrMsgTwo, err)
			}
		}
	})

	// Case dish already exists
	t.Run("AddDish_ExistingProduct", func(t *testing.T) {
		groceries := Groceries{
			Name:   "Existing Apple",
			Energy: 55.0,
		}
		existingDish := Dish{
			Name: "Existing Apple Dish",
			DishComposition: []DishContents{
				{
					DishGroceries:    &groceries,
					DishQuantityGram: 250.0,
				},
			},
		}
		dm.DishMap[existingDish.Name] = existingDish

		dish := Dish{
			Name: "Existing Apple Dish",
			DishComposition: []DishContents{
				{
					DishGroceries:    &groceries,
					DishQuantityGram: 300.0,
				},
			},
		}

		err := dm.NewDish(dish.Name, dish.DishComposition)
		expectedErrMsg := "Gericht mit diesem Namen existiert bereits"
		if err == nil || err.Error() != expectedErrMsg {
			t.Errorf("Expected error: %s, got %v", expectedErrMsg, err)
		}
	})
}

func TestEditDishAttributes(t *testing.T) {
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

	// Create a DishManager instance with some initial data
	dm := &DishManager{
		DishMap: map[string]Dish{
			"Spaghetti": {
				Name: "Spaghetti",
				DishComposition: []DishContents{
					{
						DishGroceries:    &groceriesTest,
						DishDishes:       nil,
						DishQuantityGram: 200.0,
						DishQuantityUnit: 0,
					},
				},
			},
		},
	}

	tests := []struct {
		name        string
		newName     string
		newContents []DishContents
		expectedErr error
		expectedMap map[string]Dish
	}{
		{
			name:        "Spaghetti",
			newName:     "Pasta",
			newContents: []DishContents{{DishGroceries: &groceriesTest, DishQuantityGram: 300.0, DishQuantityUnit: 0}},
			expectedErr: nil,
			expectedMap: map[string]Dish{
				"Pasta": {
					Name: "Pasta",
					DishComposition: []DishContents{
						{
							DishGroceries:    &groceriesTest,
							DishDishes:       nil,
							DishQuantityGram: 300.0,
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
		},
	}
	// Check if Name and Content edited successfully
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := dm.EditDishAttributeName(tt.name, tt.newName)
			err = dm.EditDishAttributeContent(tt.newName, tt.newContents)
			if !reflect.DeepEqual(err, tt.expectedErr) {
				t.Errorf("Got error %v, but expected %v", err, tt.expectedErr)
			}

			if !reflect.DeepEqual(dm.DishMap, tt.expectedMap) {
				t.Errorf("Got map %v, but expected %v", dm.DishMap, tt.expectedMap)
			}
		})
	}
}

func TestDeleteDish(t *testing.T) {
	// Create a DishManager instance for testing
	dm := &DishManager{
		DishMap: make(map[string]Dish),
	}

	// Add a sample dish for testing
	initialDish := Dish{
		Name: "Lasagna",
	}
	dm.DishMap[initialDish.Name] = initialDish

	// Test Case 1: Deleting a dish successfully
	err := dm.DeleteDish("Lasagna")
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	// Check if the dish is removed from the map
	_, exists := dm.DishMap["Lasagna"]
	if exists {
		t.Errorf("Expected 'Lasagna' to be removed from the map")
	}

	// Test Case 2: Deleting a non-existent dish (should fail)
	err = dm.DeleteDish("NonExistentDish")
	expectedErrMsg := "Gericht nicht gefunden"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error: %s, got %v", expectedErrMsg, err)
	}
}

func TestCalculateNutritionalValues(t *testing.T) {
	// Groceries for Tests
	apple := Groceries{
		Name:            "Apple",
		Energy:          52.0,
		Fat:             0.2,
		Carbohydrates:   13.8,
		Sugar:           10.4,
		Protein:         0.3,
		Salt:            0.0,
		Fiber:           2.4,
		Water:           85.6,
		Weight:          100.0,
		PackagingWeight: 5.0,
	}

	// Dish with groceries
	appleDishContents := DishContents{
		DishGroceries:    &apple,
		DishQuantityGram: 200.0,
		DishQuantityUnit: 0,
	}

	appleDish := Dish{
		Name:            "Apple Dish",
		DishComposition: []DishContents{appleDishContents},
	}

	// Contents of Dish
	mainDishContents := DishContents{
		DishDishes:       &appleDish,
		DishQuantityGram: 0,
		DishQuantityUnit: 1.0,
	}

	mainDish := Dish{
		Name:            "Main Dish",
		DishComposition: []DishContents{mainDishContents},
	}

	calculatedValues := mainDish.CalculateNutritionalValues()

	// Expected nutritional values for 200 grams of apples
	expectedValues := Groceries{
		Energy:        104.0,
		Fat:           0.4,
		Carbohydrates: 27.6,
		Sugar:         20.8,
		Protein:       0.6,
		Salt:          0.0,
		Fiber:         4.8,
		Water:         171.2,
	}

	// Compare calculated values with expected values
	if calculatedValues != expectedValues {
		t.Errorf("Expected: %+v, Got: %+v", expectedValues, calculatedValues)
	}
}

func TestCalculateNutritionalValues_Dish(t *testing.T) {
	// Define groceries for the first dish
	groceries := Groceries{
		Name:          "Apple",
		Energy:        52.0,
		Fat:           0.2,
		Carbohydrates: 13.8,
		Sugar:         10.4,
		Protein:       0.3,
		Salt:          0.0,
		Fiber:         2.4,
		Water:         85.6,
	}

	// Define dish contents for the first dish
	dishContents := DishContents{
		DishGroceries:    &groceries,
		DishQuantityGram: 200.0,
		DishQuantityUnit: 0,
	}

	// Create a Dish with the groceries
	dish := Dish{
		Name:            "Apple Dish",
		DishComposition: []DishContents{dishContents},
	}

	// Define the second dish with the first dish as an ingredient
	secondDishContents := DishContents{
		DishDishes:       &dish,
		DishQuantityGram: 0,
		DishQuantityUnit: 1.0,
	}

	secondDish := Dish{
		Name:            "Two Apples Dish",
		DishComposition: []DishContents{secondDishContents},
	}

	// Calculate nutritional values for the second dish
	calculatedValues := secondDish.CalculateNutritionalValues()

	// Expected nutritional values for 1 portion of the second dish
	expectedValues := Groceries{
		Energy:        104.0,
		Fat:           0.4,
		Carbohydrates: 27.6,
		Sugar:         20.8,
		Protein:       0.6,
		Salt:          0.0,
		Fiber:         4.8,
		Water:         171.2,
	}

	// Compare calculated values with expected values
	if calculatedValues != expectedValues {
		t.Errorf("Expected: %+v, Got: %+v", expectedValues, calculatedValues)
	}
}
