package data

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
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
	Price float32 `json:"price" validate:"required,gte=0"`

	// the SKU for this Product
	//
	// required: true
	// pattern: [a-z]{3}-[a-z]{3}-[\d]{3}
	SKU string `json:"sku" validate:"required,sku"`
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

func GetProductByID(id int) (*Product, error) {
	i := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}
}


func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	pl := productList[len(productList)-1]
	return pl.ID + 1
}

func UpdateProduct(prod *Product) error {
	for i, p := range productList {
		if prod.ID == p.ID {
			productList[i] = prod
			return nil
		}
	}
	return errors.New("product not found")
}

func DeleteProduct(id int) error {
	for i, p := range productList {
		if id == p.ID {
			productList = append(productList[:i], productList[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
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
