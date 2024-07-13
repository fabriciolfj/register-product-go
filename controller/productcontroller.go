package controller

import (
	"encoding/json"
	"github.com/fabriciolfj/register-product-go/entities"
	"github.com/fabriciolfj/register-product-go/service"
	"net/http"
)

type ProductController struct {
	service *service.ProductService
}

func NewProductController(s *service.ProductService) *ProductController {
	return &ProductController{
		service: s,
	}
}

func (pc *ProductController) create(w http.ResponseWriter, r *http.Request) {
	var product entities.Product
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		badRequest(w, err)
		return
	}

	err = pc.service.Save(&product)
	if err != nil {
		badRequest(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func badRequest(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
}

func (pc *ProductController) HandleProduct(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		pc.create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
