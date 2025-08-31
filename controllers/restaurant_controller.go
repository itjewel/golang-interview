package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"golang-interview/models"
	"golang-interview/service"
)

type RestaurantController struct {
	Service *service.RestaurantService
}

// GET /purchase
func (c *RestaurantController) GetHistory(w http.ResponseWriter, r *http.Request) {

	// ctx := r.Context()
	var req models.PurchaseHistory
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	data, err := c.Service.GetHistory(req)
	log.Println(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// POST /purchase/add
func (c *RestaurantController) Purchase(w http.ResponseWriter, r *http.Request) {
		//ctx := r.Context()
	var req models.PurchaseHistory
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err := c.Service.PurchaseOrder(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "purchase success",
	})
}

