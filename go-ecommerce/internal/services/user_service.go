package services

import (
	"errors"

	"go-ecommerce/internal/models"

	"golang.org/x/crypto/bcrypt"
)

// Mock database of users
var users = map[string]*models.User{}

// Register a new user
func RegisterUser(name, email, password, passwordConfirmation string) (*models.User, error) {
	mu.Lock()
	defer mu.Unlock()

	// Check if passwords match
	if password != passwordConfirmation {
		return nil, errors.New("passwords do not match")
	}

	// Check if user already exists
	if _, exists := users[email]; exists {
		return nil, errors.New("user already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create the user
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}
	users[email] = user

	return user, nil
}

// Authenticate a user
func AuthenticateUser(email, password string) (*models.User, error) {
	user, exists := users[email]
	if !exists {
		return nil, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}
	return user, nil
}
