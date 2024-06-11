//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/leoseiji/go-clean-architecture/internal/entity"
	"github.com/leoseiji/go-clean-architecture/internal/infra/database"
	"github.com/leoseiji/go-clean-architecture/internal/infra/web"
	"github.com/leoseiji/go-clean-architecture/internal/usecase"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

func NewCreateOrderUseCase(db *sql.DB) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
