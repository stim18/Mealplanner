// Matrikelnummern: 5911189 und 8441837
package data

import "meinProjekt/content"

type CurrentUser struct {
	Username string
	Password string
}

var CurrentUserInstance CurrentUser // Die Instanz der angemeldeten Benutzerdaten

func SetCurrentUser(username, password string) {
	adjustedPassword := adjustPassword(password)
	CurrentUserInstance = CurrentUser{Username: username, Password: adjustedPassword}
}

func GetCurrentUser() CurrentUser {
	return CurrentUserInstance
}

func ClearCurrentUser() {
	content.GM.GroceriesMap = make(map[string]content.Groceries)
	content.DM.DishMap = make(map[string]content.Dish)
	content.MM.MealMap = make(map[string]content.Meal)
	content.DAYM.DayMap = make(map[string]content.Day)

	CurrentUserInstance = CurrentUser{}
}

func adjustPassword(password string) string {
	targetLength := 16

	if len(password) == targetLength {
		return password
	} else if len(password) > targetLength {
		return password[:targetLength]
	} else {
		multiplier := targetLength / len(password)
		remainder := targetLength % len(password)

		adjustedPassword := password
		for i := 0; i < multiplier-1; i++ {
			adjustedPassword += password
		}
		adjustedPassword += password[:remainder]

		return adjustedPassword
	}
}
