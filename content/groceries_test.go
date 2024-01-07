// Matrikelnummern: 5911189 und 8441837
package content

import (
	"testing"
)

func TestAddGroceries(t *testing.T) {
	gm := &GroceriesManager{
		GroceriesMap: make(map[string]Groceries),
	}

	// groceries for testing
	groceries1 := Groceries{
		Name:            "TestItem0",
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
		PackagingType:   "Tüte",
		PackagingWeight: 0,
	}

	// groceries for testing
	groceries2 := Groceries{
		Name:            "",
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
		PackagingType:   "",
		PackagingWeight: 0,
	}

	//groceries with a weight of 50 grams
	groceries50Grams := Groceries{
		Name:          "TestItem1",
		Energy:        50.0,
		Fat:           2.0,
		SaturatedFats: 1.0,
		Carbohydrates: 10.0,
		Sugar:         5.0,
		Protein:       3.0,
		Salt:          0.5,
		Fiber:         1.5,
		Water:         20.0,
		Weight:        50.0,
	}

	//groceries with expected values after conversion to 100 grams
	groceries100Grams := Groceries{
		Name:          "TestItem2",
		Energy:        100.0,
		Fat:           4.0,
		SaturatedFats: 2.0,
		Carbohydrates: 20.0,
		Sugar:         10.0,
		Protein:       6.0,
		Salt:          1.0,
		Fiber:         3.0,
		Water:         40.0,
		Weight:        100.0,
	}

	//groceries with expected values after conversion to 200 grams
	groceries1Bag200Grams := Groceries{
		Name:            "TestItem3",
		Energy:          50.0,
		Fat:             2.0,
		SaturatedFats:   1.0,
		Carbohydrates:   10.0,
		Sugar:           5.0,
		Protein:         3.0,
		Salt:            0.5,
		Fiber:           1.5,
		Water:           20.0,
		PackagingType:   "Bag",
		PackagingWeight: 200,
	}

	//groceries with 200 grams and bag
	groceries1BagIn100Grams := Groceries{
		Name:            "TestItem4",
		Energy:          25.0,
		Fat:             1.0,
		SaturatedFats:   0.5,
		Carbohydrates:   5.0,
		Sugar:           2.5,
		Protein:         1.5,
		Salt:            0.25,
		Fiber:           0.75,
		Water:           10.0,
		PackagingType:   "Bag",
		PackagingWeight: 200,
	}

	// negative Input on sugar
	groceriesNegativeInput := Groceries{
		Name:            "TestItem4",
		Energy:          25.0,
		Fat:             1.0,
		SaturatedFats:   0.5,
		Carbohydrates:   5.0,
		Sugar:           -2.5,
		Protein:         1.5,
		Salt:            0.25,
		Fiber:           0.75,
		Water:           10.0,
		PackagingType:   "Bag",
		PackagingWeight: 200,
	}

	// Test Case 1: Adding new groceries successfully
	err := gm.NewGroceries(groceries1)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	// Test Case 2: Adding duplicate groceries (should fail)
	err = gm.NewGroceries(groceries1)
	expectedErrMsg := "Lebensmittel mit diesem Namen existiert bereits"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error: %s, got %v", expectedErrMsg, err)
	}

	// Test Case 3: Adding groceries without a name(should fail)
	err = gm.NewGroceries(groceries2)
	expectedErrMsg = "Ungültiger Eingabewert: Lebensmittelname darf nicht leer sein"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error: %s, got %v", expectedErrMsg, err)
	}

	// Test Case 4: Lebensmittel mit 50 Gramm hinzufügen und auf 100 Gramm umrechnen
	err = gm.NewGroceries(groceries50Grams)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	// Test Case 5: editing all attributes to new gram
	updatedValues := gm.GroceriesMap["TestItem1"]
	if updatedValues.Energy != groceries100Grams.Energy {
		t.Errorf("Expected updated values: %+v, got %+v", groceries100Grams.Energy, updatedValues.Energy)
	}

	// Test Case 6: Add new groceries of 50 grams and convert to 100 grams
	err = gm.NewGroceries(groceries1Bag200Grams)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	// Test Case 7: CHeck if the groceries got converted to 100 grams
	updatedValues = gm.GroceriesMap["TestItem3"]
	if updatedValues.Energy != groceries1BagIn100Grams.Energy {
		t.Errorf("Expected updated values: %+v, got %+v", groceries1BagIn100Grams.Energy, updatedValues.Energy)
	}

	// Test Case 8: negative Input on one Attribute
	err = gm.NewGroceries(groceriesNegativeInput)
	expectedErrMsg = "Fehlerhafte Eingabe: Die Attributwerte dürfen nicht kleiner als 0 sein"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error: %s, got %v", expectedErrMsg, err)
	}
}

/*
func TestEditGroceries(t *testing.T) {
	gm := &GroceriesManager{
		GroceriesMap: make(map[string]Groceries),
	}

	// Preparing initial groceries entry
	initialGroceries := Groceries{
		Name:          "Milch",
		Energy:        42.0,
		Fat:           3.4,
		SaturatedFats: 2.0,
		Carbohydrates: 4.7,
		Sugar:         4.7,
		Protein:       3.4,
		Salt:          0.1,
		Fiber:         0.0,
		Water:         87.5,
		Weight:        250.0,
		PackagingType: "Tüte",
	}

	// Preparing initial groceries entry
	initialGroceries1 := Groceries{
		Name:          "Milch",
		Energy:        55.0,
		Fat:           3.4,
		SaturatedFats: 2.0,
		Carbohydrates: 4.7,
		Sugar:         4.7,
		Protein:       3.4,
		Salt:          0.1,
		Fiber:         0.0,
		Water:         87.5,
		Weight:        250.0,
		PackagingType: "Tüte",
	}

	// Preparing initial groceries entry
	initialGroceries2 := Groceries{
		Name:          "Milch",
		Energy:        42.0,
		Fat:           3.4,
		SaturatedFats: 2.0,
		Carbohydrates: 4.7,
		Sugar:         -4.7,
		Protein:       3.4,
		Salt:          0.1,
		Fiber:         0.0,
		Water:         87.5,
		Weight:        250.0,
		PackagingType: "Tüte",
	}

	err := gm.NewGroceries(initialGroceries)
	if err != nil {
		t.Errorf("Anlegen hat nicht geklappt")
	}
	// Test Case 1: Editing an existing attribute successfully
	err = gm.EditGroceriesAttribute(initialGroceries1)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	updatedValue := gm.GroceriesMap["Milch"].Energy
	expectedValue := 55.0
	if updatedValue != expectedValue {
		t.Errorf("Expected updated value: %f, got %f", expectedValue, updatedValue)
	}

	// check, if other values stay the same
	unchangedValue := gm.GroceriesMap["Milch"].Sugar
	unchangedExpectedValue := 4.7
	if unchangedValue != unchangedExpectedValue {
		t.Errorf("Expected value: %f, got %f", unchangedExpectedValue, unchangedValue)
	}

	// Test Case 2: Editing with a negative new Value
	err = gm.EditGroceriesAttribute(initialGroceries2)
	expectedErrMsg := "Fehlerhafte Eingabe: Die Attributwerte dürfen nicht kleiner als 0 sein"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error: %s, got %v", expectedErrMsg, err)
	}

}
*/

func TestDeleteGroceries(t *testing.T) {
	gm := &GroceriesManager{
		GroceriesMap: make(map[string]Groceries),
	}

	// Preparing initial groceries entry
	initialGroceries := Groceries{
		Name:            "Banana",
		Energy:          89.0,
		Fat:             0.3,
		SaturatedFats:   0.1,
		Carbohydrates:   22.8,
		Sugar:           12.2,
		Protein:         1.1,
		Salt:            0.01,
		Fiber:           2.6,
		Water:           74.91,
		Weight:          100.0,
		PackagingWeight: 5.0,
	}
	gm.GroceriesMap[initialGroceries.Name] = initialGroceries

	// Test Case 1: Deleting existing groceries successfully
	err := gm.DeleteGroceries("Banana")
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	_, exists := gm.GroceriesMap["Banana"]
	if exists {
		t.Errorf("Expected 'Banana' groceries to be deleted")
	}

	// Test Case 2: Deleting non-existent groceries (should fail)
	err = gm.DeleteGroceries("NonExistentGroceries")
	expectedErrMsg := "Lebensmittel nicht gefunden"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error: %s, got %v", expectedErrMsg, err)
	}

	//Test Case 3: Deleting without input (should fail)
	err = gm.DeleteGroceries("")
	expectedErrMsg = "Ungültiger Eingabewert: Lebensmittelname darf nicht leer sein"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error: %s, got %v", expectedErrMsg, err)
	}
}
