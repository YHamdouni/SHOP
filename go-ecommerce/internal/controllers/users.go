package controllers

import (
	"html/template"
	"log"
	"net/http"

	"go-ecommerce/internal/services"

	"github.com/gorilla/sessions"
)

// Initialize session store with a secure key
var store = sessions.NewCookieStore([]byte("your-secret-key"))

// RegisterHandler handles user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		passwordConfirmation := r.FormValue("password_confirmation")

		// Register user with provided details
		_, err := services.RegisterUser(name, email, password, passwordConfirmation)
		if err != nil {
			// Provide feedback to the user if registration fails
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Redirect to login page after successful registration
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Render registration page for GET requests
	tmpl, err := template.ParseFiles("web/templates/register.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// LoginHandler handles user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Authenticate user with provided credentials
		user, err := services.AuthenticateUser(email, password)
		if err != nil {
			// Log failed login attempts and provide feedback to the user
			log.Printf("Login failed for email %s: %v", email, err)
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// Create a new session for the authenticated user
		session, _ := store.Get(r, "session-name")
		session.Values["user"] = user.Email
		err = session.Save(r, w)
		if err != nil {
			log.Printf("Error saving session: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Redirect to products page after successful login
		http.Redirect(w, r, "/products", http.StatusSeeOther)
		return
	}

	// Render login page for GET requests
	tmpl, err := template.ParseFiles("web/templates/login.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
