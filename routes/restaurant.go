package routes

import (
	"net/http"

	"golang-interview/controllers"
	"golang-interview/repository"
	"golang-interview/service"
)

func RestaurantRoutes(mux *http.ServeMux) {
	restaurantRepo := &repository.RestaurantRepository{}
	restaurantService := &service.RestaurantService{Repo: restaurantRepo}
	restaurantController := &controllers.RestaurantController{Service: restaurantService}

	mux.HandleFunc("POST /restaurant/purchase", restaurantController.Purchase) // Pharchase
	mux.HandleFunc("POST /restaurant/history", restaurantController.GetHistory) // GET
}
