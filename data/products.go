package data

import (
	"encoding/json"
	"errors"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"` // >0
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	rex := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := rex.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	pl := productList[len(productList)-1]
	return pl.ID + 1
}

func UpdateProduct(id int, prod *Product) error {
	for i, p := range productList {
		if id == p.ID {
			prod.ID = id
			productList[i] = prod
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
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and string coffee without milk",
		Price:       1.99,
		SKU:         "zcv323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
