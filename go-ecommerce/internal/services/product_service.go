package services

import (
	"errors"
	"sync"

	"go-ecommerce/internal/models"
)

// In-memory store for products
var (
	products = make(map[int]*models.Product)
	nextID   = 1
	mu       sync.Mutex
)

// CreateProduct adds a new product to the store
func CreateProduct(title, description, brand, model, condition, color string, price float64) (*models.Product, error) {
	mu.Lock()
	defer mu.Unlock()

	// Create a new product
	product := &models.Product{
		ID:          nextID,
		Title:       title,
		Description: description,
		Brand:       brand,
		Model:       model,
		Condition:   condition,
		Color:       color,
		Price:       price,
	}

	// Add the product to the store
	products[nextID] = product
	nextID++

	return product, nil
}

// GetProduct retrieves a product by its ID
func GetProduct(id int) (*models.Product, error) {
	mu.Lock()
	defer mu.Unlock()

	product, exists := products[id]
	if !exists {
		return nil, errors.New("product not found")
	}
	return product, nil
}

// ListProducts returns a slice of all products
func ListProducts() []*models.Product {
	mu.Lock()
	defer mu.Unlock()

	// Create a slice to hold all products
	productList := make([]*models.Product, 0, len(products))
	for _, product := range products {
		productList = append(productList, product)
	}

	return productList
}

// UpdateProduct updates the details of an existing product
func UpdateProduct(id int, title, description, brand, model, condition, color string, price float64) (*models.Product, error) {
	mu.Lock()
	defer mu.Unlock()

	product, exists := products[id]
	if !exists {
		return nil, errors.New("product not found")
	}

	// Update product fields
	product.Title = title
	product.Description = description
	product.Brand = brand
	product.Model = model
	product.Condition = condition
	product.Color = color
	product.Price = price

	return product, nil
}

// DeleteProduct removes a product from the store
func DeleteProduct(id int) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := products[id]; !exists {
		return errors.New("product not found")
	}

	delete(products, id)
	return nil
}
