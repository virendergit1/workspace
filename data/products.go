package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	// use struct tag to format json
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"desc"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// for ease of use define custom type for list of products i.e []*Product
type Products []*Product

// method to convert a Product to a JSON object and write out
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// method to convert a JSON to a Product
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProducts() Products {
	return productsList
}

// add the Product to the existing list
func AddProduct(p *Product) {
	p.ID = len(productsList) + 1
	productsList = append(productsList, p)
}

var productsList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky cofffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	}, {
		ID:          2,
		Name:        "Espresso",
		Description: "Shot and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
