package presenter

import (
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
