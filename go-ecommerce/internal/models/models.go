package models

type User struct {
    ID       int
    Name     string
    Email    string
    Password string
}

type Product struct {
    ID          int
    Title       string
    Price       float64
    Description string
    Brand       string
    Model       string
    Condition   string
    Color       string
}

type CartItem struct {
    ProductID int
    Quantity  int
}
