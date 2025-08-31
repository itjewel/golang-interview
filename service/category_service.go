package service

import (
	"context"
	"errors"
	"golang-interview/models"
	"golang-interview/repository"
)

type CategoryService struct {
	Repo *repository.CategoryRepository
}

// Get all categories
func (s *CategoryService) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	return s.Repo.GetAll(ctx)
}

// Get category by ID
func (s *CategoryService) GetCategoryByID(id int) (*models.Category, error) {
	return s.Repo.GetByID(id)
}

// Add category
func (s *CategoryService) AddCategory(c models.Category) (*models.Category, error) {
	if c.Name == "" {
		return nil, errors.New("category name cannot be empty")
	}
	id, err := s.Repo.Create(c)
	if err != nil {
		return nil, err
	}
	c.ID = int(id)
	return &c, nil
}

// Update category
func (s *CategoryService) UpdateCategory(c models.Category) error {
	if c.ID == 0 {
		return errors.New("invalid category ID")
	}
	rowsAffected, err := s.Repo.Update(c)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no category found to update")
	}
	return nil
}

// Delete category
func (s *CategoryService) DeleteCategory(id int) error {
	rowsAffected, err := s.Repo.Delete(id)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no category found to delete")
	}
	return nil
}

// Search by name
func (s *CategoryService) SearchCategoryByName(name string) ([]models.Category, error) {
	return s.Repo.SearchByName(name)
}

// Get by price range
func (s *CategoryService) GetCategoriesByPriceRange(from, to float64) ([]models.Category, error) {
	return s.Repo.GetByPriceRange(from, to)
}

func (s *CategoryService) BulkUpload(ctx context.Context, getValue []models.Category) (string, error) {
	value, err := s.Repo.Seeding(ctx, getValue)
	if err != nil {
		return "Not uploded", err
	}
	return value, nil
}
