package entities

import uuid "github.com/satori/go.uuid"

type Product struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Describe string `gorm:"size:255" json:"describe"`
	Uuid     string `gorm:"size:255" json:"uuid"`
}

func (Product) TableName() string { return "products" }

func (p *Product) GnerateUUID() {
	p.Uuid = uuid.NewV4().String()
}
