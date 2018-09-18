package model

import (
	"github.com/graphql-go/graphql"
)

type Store struct {
	ID       int       `json:"store_id" required:"true"`
	Name     string    `json:"store_name" required:"true"`
	Products []Product `json:"products,omitempty"`
}

var StoreGraphqlObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Store",
	Fields:      graphql.BindFields(Store{}),
	Description: "Store object model",
})

func (s *Store) MakeFields() graphql.Fields {
	storeField := make(graphql.Fields)

	storeField["store_id"] = &graphql.Field{
		Name: "StoreID",
		Type: graphql.Int,
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			if store, ok := p.Source.(*Store); ok {
				i = store.ID
			} else {
				i = p.Source.(Store).ID
			}
			return
		},
	}

	storeField["store_name"] = &graphql.Field{
		Name: "StoreName",
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			if store, ok := p.Source.(*Store); ok {
				i = store.Name
			} else {
				i = p.Source.(Store).Name
			}
			return
		},
	}

	return storeField
}
