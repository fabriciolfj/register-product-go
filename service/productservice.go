package service

import (
	"github.com/fabriciolfj/register-product-go/entities"
	"github.com/fabriciolfj/register-product-go/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repository,
	}
}

func (ps *ProductService) Save(p *entities.Product) error {
	p.GnerateUUID()
	err := ps.repo.Save(p)
	if err != nil {
		return err
	}

	return nil
}
