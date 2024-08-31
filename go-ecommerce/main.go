package main

import (
	"net/http"

	"go-ecommerce/internal/controllers"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key")) // Use a more secure key in production

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/index.html")
	})
	r.HandleFunc("/register", controllers.RegisterHandler)
	r.HandleFunc("/login", controllers.LoginHandler)
	r.HandleFunc("/products", controllers.CreateProductHandler)
	// Add more routes here

	fs := http.FileServer(http.Dir("web/static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", r)
}
