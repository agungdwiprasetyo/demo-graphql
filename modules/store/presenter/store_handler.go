package presenter

import (
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/usecase"
	"github.com/labstack/echo"
)

type StoreHandler struct {
	storeUsecase usecase.StoreUsecase
}

func NewStoreHandler(storeUsecase usecase.StoreUsecase) *StoreHandler {
	return &StoreHandler{storeUsecase}
}

func (h *StoreHandler) MountGraphQL(router *echo.Group) {
	router.GET("/graphql", h.InitGraphQL)
}

func (h *StoreHandler) MountREST(router *echo.Group) {
	router.POST("/add", h.SaveStore)
}
