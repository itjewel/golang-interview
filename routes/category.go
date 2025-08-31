package routes

import (
	"net/http"

	"golang-interview/controllers"
	"golang-interview/repository"
	"golang-interview/service"
)

func CategoryRoutes(mux *http.ServeMux) {
	categoryRepo := &repository.CategoryRepository{}
	categoryService := &service.CategoryService{Repo: categoryRepo}
	categoryController := &controllers.CategoryController{Service: categoryService}

	mux.HandleFunc("GET /categories", categoryController.GetCategories) // GET
	mux.HandleFunc("GET /category/bulk-upload", categoryController.BulkUpload)
}
