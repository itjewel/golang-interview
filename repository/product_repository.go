package repository

import (
	"context"
	"golang-interview/database"
	"golang-interview/models"
)

type CategoryRepository struct{}

// Get all categories
func (r *CategoryRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	rows, err := database.DB.QueryContext(ctx, "SELECT id, name, price FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Price); err != nil {
			continue
		}
		categories = append(categories, c)
	}
	return categories, nil
}

// Get category by ID
func (r *CategoryRepository) GetByID(id int) (*models.Category, error) {
	var c models.Category
	err := database.DB.QueryRow("SELECT id, name, price FROM categories WHERE id = ?", id).Scan(&c.ID, &c.Name, &c.Price)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// Insert category
func (r *CategoryRepository) Create(c models.Category) (int64, error) {
	res, err := database.DB.Exec("INSERT INTO categories (name, price) VALUES (?, ?)", c.Name, c.Price)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// Update category
func (r *CategoryRepository) Update(c models.Category) (int64, error) {
	res, err := database.DB.Exec("UPDATE categories SET name=?, price=? WHERE id=?", c.Name, c.Price, c.ID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// Delete category
func (r *CategoryRepository) Delete(id int) (int64, error) {
	res, err := database.DB.Exec("DELETE FROM categories WHERE id=?", id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// Search by name
func (r *CategoryRepository) SearchByName(name string) ([]models.Category, error) {
	rows, err := database.DB.Query("SELECT id,name FROM categories WHERE name LIKE ?", "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			continue
		}
		categories = append(categories, c)
	}
	return categories, nil
}

// Range query by price
func (r *CategoryRepository) GetByPriceRange(from, to float64) ([]models.Category, error) {
	rows, err := database.DB.Query("SELECT id,name,price FROM categories WHERE price BETWEEN ? AND ?", from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Price); err != nil {
			continue
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (r *CategoryRepository) Seeding(ctx context.Context, getValue []models.Category) (string, error) {
	for _, value := range getValue {
		_, err := database.DB.ExecContext(ctx, "INSERT INTO categories (name,price) VALUES (?,?)", value.Name, value.Price)
		if err != nil {
			// return "not update", nil
			continue
		}

	}
	return "insert Succfully", nil
}
