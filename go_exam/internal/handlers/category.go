package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"go_exam/internal/models"
	"go_exam/internal/repository"
)

type CategoryPageData struct {
	Category *models.Category
}

func CategoryPage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/category/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		repo := repository.NewCategoryRepository(db)
		category, err := repo.GetByID(id)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		tmpl := template.Must(template.ParseFiles("templates/category.html"))
		tmpl.Execute(w, CategoryPageData{
			Category: category,
		})
	}
}
