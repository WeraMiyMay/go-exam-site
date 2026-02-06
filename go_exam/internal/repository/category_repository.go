package repository

import (
	"database/sql"
	"go_exam/internal/models"
)

type CategoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) GetAll() ([]models.Category, error) {
	rows, err := r.DB.Query(`
		SELECT id, name, description
		FROM categories
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category

	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}
