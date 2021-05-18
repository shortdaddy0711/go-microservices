package data

import (
	"fmt"
)

// ErrProductNotFound is an error raised when a product can not be found in the DB
var ErrProductNotFound = fmt.Errorf("Product not found")

// Product defines the structure for an product API
// swagger:model
type Product struct {
	// the id for the Product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this product
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this product
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for this Product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`

	// the SKU for this Product
	//
	// required: true
	// pattern: [a-z]{3}-[a-z]{3}-[\d]{3}
	SKU string `json:"sku" validate:"sku"`
	// CreatedOn   string  `json:"-"`
	// UpdatedOn   string  `json:"-"`
	// DeletedOn   string  `json:"-"`
}

// Products defines a slice of Product
type Products []*Product

// findIndexByProductID finds the index of a product in the DB
// returns -1 if there is no matching product
func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}
	return -1
}

// GetProducts returns all products from the DB
func GetProducts() Products {
	return productList
}

// GetProductByID returns a single product which matches the id from the DB
// and returns a ProductNotFound error if no match in the DB
func GetProductByID(id int) (*Product, error) {
	i := findIndexByProductID(id)
	if i == -1 {
		return nil, ErrProductNotFound
	}
	return productList[i], nil
}

// AddProduct adds a new product to the DB
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

// getNextID returns id number for new product
func getNextID() int {
	pl := productList[len(productList)-1]
	return pl.ID + 1
}

// UpdateProduct replaces a product in the DB with the given item
// and returns a ProductNotFound error if no match in the DB
func UpdateProduct(p Product) error {
	i := findIndexByProductID(p.ID)
	if i == -1 {
		return ErrProductNotFound
	}

	// update the product
	productList[i] = &p

	return nil
}

// DeleteProduct deletes a product in the DB
// and returns a ProductNotFound error if no match in the DB
func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1:]...)

	return nil
}

var productList = Products{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		// CreatedOn:   time.Now().UTC().String(),
		// UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and string coffee without milk",
		Price:       1.99,
		SKU:         "zcv323",
		// CreatedOn:   time.Now().UTC().String(),
		// UpdatedOn:   time.Now().UTC().String(),
	},
}
