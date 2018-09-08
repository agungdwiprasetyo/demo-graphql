package presenter

import (
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/model"
	"github.com/agungdwiprasetyo/go-utils/debug"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
)

func (h *StoreHandler) GetAllStore(c echo.Context) error {
	query := c.QueryParam("query")
	debug.Println("query:", query)

	schema, err := h.storeUsecase.GetStore()
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	return c.JSON(200, result)
}

func (h *StoreHandler) SaveStore(c echo.Context) error {
	var payload model.Store
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	debug.PrintJSON(payload)
	// return nil

	err := h.storeUsecase.SaveStore(&payload)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, "Success")
}
