package usecase

import (
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/model"
	"github.com/graphql-go/graphql"
)

func (uc *storeUsecase) GetAllStore() *graphql.Field {
	stores := new(model.Store).MakeFields()
	stores["products"] = &graphql.Field{
		Name: "Products",
		Type: graphql.NewList(model.ProductGraphqlObject),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			store, _ := p.Source.(model.Store)
			return uc.productQuery.FindByStoreID(store.ID)
		},
	}

	return &graphql.Field{
		Name: "stores",
		Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
			Name:   "Stores",
			Fields: stores,
		})),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return uc.storeQuery.FindAll()
		},
	}
}

func (uc *storeUsecase) GetStoreByID() *graphql.Field {
	store := model.StoreGraphqlObject
	store.AddFieldConfig("products", &graphql.Field{
		Name: "Products",
		Type: graphql.NewList(model.ProductGraphqlObject),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			store, _ := p.Source.(*model.Store)
			return uc.productQuery.FindByStoreID(store.ID)
		},
	})

	return &graphql.Field{
		Name: "store",
		Type: store,
		Args: graphql.FieldConfigArgument{
			"store_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, _ := p.Args["store_id"].(int)
			return uc.storeQuery.FindByID(id)
		},
	}
}

func (uc *storeUsecase) GetAllProduct() *graphql.Field {
	products := new(model.Product).MakeFields()
	products["store"] = &graphql.Field{
		Name: "Store",
		Type: model.StoreGraphqlObject,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			product, _ := p.Source.(model.Product)
			return uc.storeQuery.FindByID(product.StoreID)
		},
	}

	return &graphql.Field{
		Name: "products",
		Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
			Name:   "Products",
			Fields: products,
		})),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return uc.productQuery.FindAll()
		},
	}
}

func (uc *storeUsecase) GetProductByID() *graphql.Field {
	product := model.ProductGraphqlObject
	product.AddFieldConfig("store", &graphql.Field{
		Name: "Store",
		Type: model.StoreGraphqlObject,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			product, _ := p.Source.(*model.Product)
			return uc.storeQuery.FindByID(product.StoreID)
		},
	})

	return &graphql.Field{
		Name: "product",
		Type: product,
		Args: graphql.FieldConfigArgument{
			"product_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, _ := p.Args["product_id"].(int)
			return uc.productQuery.FindByID(id)
		},
	}
}
