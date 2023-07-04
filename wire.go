//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/eltoncasacio/di-wire/product"
	"github.com/google/wire"
)

var setRepositoryDependencies = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewProductUsecase(db *sql.DB) *product.ProductUsecase {
	wire.Build(
		setRepositoryDependencies,
		product.NewProductUsecase,
	)
	return &product.ProductUsecase{}
}
