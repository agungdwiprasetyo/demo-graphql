package main

import (
	"fmt"
	"log"

	"github.com/agungdwiprasetyo/demo-graphql/config"
	"github.com/agungdwiprasetyo/demo-graphql/middleware"
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/presenter"
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/query"
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/repository"
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/usecase"
	"github.com/labstack/echo"
)

type Service struct {
	conf         config.Config
	StoreService *presenter.StoreHandler
}

func NewService(conf config.Config) *Service {
	queryDecorator := query.NewQuery(conf.LoadReadDB())
	repositoryDecorator := repository.NewRepository(conf.LoadWriteDB())

	storeUsecase := usecase.NewStoreUsecase(queryDecorator, repositoryDecorator)

	service := new(Service)
	service.conf = conf
	service.StoreService = presenter.NewStoreHandler(storeUsecase)
	return service
}

func (serv *Service) ServeHTTP(port int) {
	app := echo.New()
	app.Use(middleware.SetCORS())

	storeGroup := app.Group("/store")
	serv.StoreService.MountGraphQL(storeGroup)
	serv.StoreService.MountREST(storeGroup)

	appPort := fmt.Sprintf(":%d", port)
	if err := app.Start(appPort); err != nil {
		log.Fatal(err)
	}
}

func (serv *Service) ServeGRPC(port int) {
	// incoming
}
