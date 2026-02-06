package repository

import (
	"database/sql"
	"go_exam/internal/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetByID(id int) (*models.Category, error) {
	row := r.db.QueryRow(`
		SELECT id, name, description
		FROM categories
		WHERE id = ?
	`, id)

	var c models.Category
	err := row.Scan(&c.ID, &c.Name, &c.Description)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *CategoryRepository) GetAll() ([]*models.Category, error) {
	rows, err := r.db.Query(`
		SELECT id, name, description
		FROM categories
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.Category

	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
			return nil, err
		}
		categories = append(categories, &c)
	}

	return categories, nil
}
