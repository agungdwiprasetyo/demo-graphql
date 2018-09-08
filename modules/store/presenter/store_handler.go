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

func (h *StoreHandler) Mount(router *echo.Group) {
	router.GET("", h.GetAllStore)
	router.POST("", h.SaveStore)
}
