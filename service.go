package main

import (
	"fmt"

	"github.com/agungdwiprasetyo/demo-graphql/modules/store/presenter"
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/query"
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/usecase"
	"github.com/labstack/echo"
)

type Service struct {
	StoreHandler *presenter.StoreHandler
}

func NewService() *Service {
	storeQuery := query.NewStoreQuery()
	productQuery := query.NewProductQuery()

	uc := usecase.NewStoreUsecase(storeQuery, productQuery)

	service := new(Service)
	service.StoreHandler = presenter.NewStoreHandler(uc)
	return service
}

func (serv *Service) ServeHTTP(port int) {
	app := echo.New()

	storeGroup := app.Group("/graphql/store")
	serv.StoreHandler.Mount(storeGroup)

	appPort := fmt.Sprintf(":%d", port)
	app.Logger.Fatal(app.Start(appPort))
}
