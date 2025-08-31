package service

import (
	"errors"
	"golang-interview/models"
	"golang-interview/repository"
)

type RestaurantService struct {
	Repo *repository.RestaurantRepository
}

// Get all Purchse History
func (s *RestaurantService) GetHistory(req models.PurchaseHistory) ([]models.PurchaseHistory, error) {
	return s.Repo.GetAllPurchaseHistory(req)
}


// Purchse restaurant
func (s *RestaurantService) PurchaseOrder(req models.PurchaseHistory) (*models.PurchaseHistory, error) {
	if req.Price == 0 {
		return nil, errors.New("purchase price cannot be empty")
	}
	id, err := s.Repo.PurchaseOrder(req)
	if err != nil {
		return nil, err
	}
	req.ID = int(id)
	return &req, nil
}

