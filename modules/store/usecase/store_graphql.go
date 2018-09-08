package usecase

import (
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/model"
	"github.com/graphql-go/graphql"
)

func (uc *storeUsecase) GetAllStore() *graphql.Field {
	storeField := make(graphql.Fields)
	storeField["store_id"] = &graphql.Field{
		Name: "StoreID",
		Type: graphql.Int,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			store := p.Source.(model.Store)
			return store.ID, nil
		},
	}

	storeField["store_name"] = &graphql.Field{
		Name: "StoreName",
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			store := p.Source.(model.Store)
			return store.Name, nil
		},
	}

	storeField["products"] = &graphql.Field{
		Name: "Products",
		Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
			Name:   "Products",
			Fields: graphql.BindFields(model.Product{}),
		})),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			store := p.Source.(model.Store)
			return uc.productQuery.GetByStoreID(store.ID)
		},
	}

	return &graphql.Field{
		Name: "stores",
		Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
			Name:   "Stores",
			Fields: storeField,
		})),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return uc.storeQuery.FindAll()
		},
	}
}

func (uc *storeUsecase) GetStoreByID() *graphql.Field {
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
			store, _ := p.Source.(*model.Store)
			return uc.productQuery.GetByStoreID(store.ID)
		},
	}

	return &graphql.Field{
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
}
