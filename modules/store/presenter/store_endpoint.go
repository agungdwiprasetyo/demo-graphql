package presenter

import (
	"net/http"

	"github.com/agungdwiprasetyo/demo-graphql/modules/store/model"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
)

// InitGraphQL endpoint -> GET /graphql/store?query={your_graphql_query}
func (h *StoreHandler) InitGraphQL(c echo.Context) error {
	query := c.QueryParam("query")
	// debug.Println("query:", query)

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"get_all_stores":  h.storeUsecase.GetAllStore(),
				"get_store":       h.storeUsecase.GetStoreByID(),
				"get_all_product": h.storeUsecase.GetAllProduct(),
				"get_product":     h.storeUsecase.GetProductByID(),
			},
		}),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if result.HasErrors() {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": result.Errors})
	}

	return c.JSON(http.StatusOK, result)
}

// SaveStore endpoint -> POST /graphql/store
func (h *StoreHandler) SaveStore(c echo.Context) error {
	var payload model.Store
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.storeUsecase.SaveStore(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
}
