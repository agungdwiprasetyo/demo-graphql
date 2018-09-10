package main

import (
	"fmt"

	"github.com/agungdwiprasetyo/demo-graphql/config"
	"github.com/agungdwiprasetyo/demo-graphql/middleware"
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/presenter"
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/query"
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/repository"
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/usecase"
	"github.com/labstack/echo"
)

type Service struct {
	StoreHandler *presenter.StoreHandler
}

func NewService() *Service {
	db := config.GetPostgresConnection()
	read := query.NewQuery(db)
	write := repository.NewRepository(db)

	uc := usecase.NewStoreUsecase(read, write)

	service := new(Service)
	service.StoreHandler = presenter.NewStoreHandler(uc)
	return service
}

func (serv *Service) ServeHTTP(port int) {
	app := echo.New()

	app.Use(middleware.SetCORS())

	storeGroup := app.Group("/graphql/store")
	serv.StoreHandler.Mount(storeGroup)

	appPort := fmt.Sprintf(":%d", port)
	app.Logger.Fatal(app.Start(appPort))
}
