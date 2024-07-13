package repository

import (
	"github.com/fabriciolfj/register-product-go/entities"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(p *entities.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) Save(p *entities.Product) error {
	return r.db.Create(&p).Error
}
