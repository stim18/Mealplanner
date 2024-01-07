// Matrikelnummern: 5911189 und 8441837
package server

import (
	"fmt"
	"log"
	"meinProjekt/handlers"
	"net/http"
)

func RunServer() {
	http.HandleFunc("/", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/currentuser", handlers.ShowCurrentUserHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)

	http.HandleFunc("/groceries", handlers.GroceriesHandler)
	http.HandleFunc("/getGroceries", handlers.GetGroceries)
	http.HandleFunc("/addGroceries", handlers.AddGroceriesHandler)
	http.HandleFunc("/editGroceries", handlers.EditGroceriesHandler)
	http.HandleFunc("/deleteGroceries", handlers.DeleteGroceriesHandler)

	http.HandleFunc("/dish", handlers.DishHandler)
	http.HandleFunc("/getDish", handlers.GetDish)
	http.HandleFunc("/deleteDish", handlers.DeleteDishHandler)
	http.HandleFunc("/addDish", handlers.AddDishHandler)
	http.HandleFunc("/editDish", handlers.EditDishHandler)

	http.HandleFunc("/meal", handlers.MealHandler)
	http.HandleFunc("/getMeal", handlers.GetMealHandler)
	http.HandleFunc("/deleteMeal", handlers.DeleteMealHandler)
	http.HandleFunc("/addMeal", handlers.AddMealHandler)
	http.HandleFunc("/editMeal", handlers.EditMealHandler)

	http.HandleFunc("/day", handlers.DayHandler)
	http.HandleFunc("/getDay", handlers.GetDay)

	fmt.Println("Server is running on https://localhost:4443")
	log.Fatal(http.ListenAndServeTLS(":4443", "certificate/cert.pem", "certificate/key.pem", nil))
}
