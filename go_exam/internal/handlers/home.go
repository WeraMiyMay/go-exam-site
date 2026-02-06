package handlers

import (
	"html/template"
	"net/http"

	"go_exam/internal/repository"
)

type HomeHandler struct {
	CategoryRepo *repository.CategoryRepository
	Tmpl         *template.Template
}

func NewHomeHandler(repo *repository.CategoryRepository) *HomeHandler {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	return &HomeHandler{
		CategoryRepo: repo,
		Tmpl:         tmpl,
	}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	categories, err := h.CategoryRepo.GetAll()
	if err != nil {
		http.Error(w, "Ошибка загрузки категорий", http.StatusInternalServerError)
		return
	}

	h.Tmpl.Execute(w, categories)
}
