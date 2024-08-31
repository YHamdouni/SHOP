package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"go-ecommerce/internal/services"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		description := r.FormValue("description")
		brand := r.FormValue("brand")
		model := r.FormValue("model")
		condition := r.FormValue("condition")
		color := r.FormValue("color")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)

		_, err := services.CreateProduct(title, description, brand, model, condition, color, price)
		if err != nil {
			http.Error(w, "Failed to create product", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/products.html")
	if err != nil {
		http.Error(w, "Internal Server Error5", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
