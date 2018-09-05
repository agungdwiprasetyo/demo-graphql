package usecase

import (
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/model"
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/query"
	"github.com/graphql-go/graphql"
)

type storeUsecase struct {
	storeQuery   query.Store
	productQuery query.Product
}

func NewStoreUsecase(storeQuery query.Store, productQuery query.Product) StoreUsecase {
	return &storeUsecase{storeQuery, productQuery}
}

func (uc *storeUsecase) GetStore() (graphql.Schema, error) {
	storeField := make(graphql.Fields)
	storeField["store_id"] = &graphql.Field{
		Name: "StoreID",
		Type: graphql.Int,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			store := p.Source.(*model.Store)
			return store.ID, nil
		},
	}

	storeField["store_name"] = &graphql.Field{
		Name: "StoreName",
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			store := p.Source.(*model.Store)
			return store.Name, nil
		},
	}

	storeField["products"] = &graphql.Field{
		Name: "Products",
		Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
			Name:   "Product",
			Fields: graphql.BindFields(model.Product{}),
		})),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			storeID := p.Source.(*model.Store).ID
			return uc.productQuery.GetByStoreID(storeID)
		},
	}

	responseData := &graphql.Field{
		Name: "store",
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Store",
			Fields: storeField,
		}),
		Args: graphql.FieldConfigArgument{
			"store_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, _ := p.Args["store_id"].(int)
			return uc.storeQuery.GetByID(id)
		},
	}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"store": responseData,
			},
		}),
	})

	return schema, err
}
