//go:build wireinject
// +build wireinject

package main

import (
	"github.com/fabriciolfj/register-product-go/controller"
	"github.com/fabriciolfj/register-product-go/datasource"
	"github.com/fabriciolfj/register-product-go/repository"
	"github.com/fabriciolfj/register-product-go/service"
	"github.com/google/wire"
)

func InitializeProductControllerWire() (*controller.ProductController, error) {
	wire.Build(
		datasource.NewDataBase,
		repository.NewProductRepository,
		service.NewProductService,
		controller.NewProductController,
	)
	return &controller.ProductController{}, nil
}
