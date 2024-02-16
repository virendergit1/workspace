package handlers

import (
	"log"
	"net/http"
	"workspace/data"
)

type ProductsHandler struct {
	l *log.Logger
}

// Idiomatic. this is like constructor which works on a Struct defined above
func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{l}
}

// This is main handler methd to handle HTTP requests
func (ph *ProductsHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		ph.getProducts(rw, r)
		return
	} else if r.Method == http.MethodPost {
		ph.addProduct(rw, r)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (ph *ProductsHandler) getProducts(rw http.ResponseWriter, r *http.Request) {
	ph.l.Println("Inside getProducts")
	products := data.GetProducts()
	// this will return the response in Json format
	err := products.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Error while parsing products", http.StatusInternalServerError)
	}
}

func (ph *ProductsHandler) addProduct(rw http.ResponseWriter, r *http.Request) {
	ph.l.Println("Inside addProduct")
	p := &data.Product{}
	// this will save the response in Json format
	err := p.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Error while parsing products", http.StatusInternalServerError)
	}

	data.AddProduct(p)
}
