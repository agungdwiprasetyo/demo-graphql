package model

import (
	"github.com/graphql-go/graphql"
)

type Product struct {
	ID      int    `json:"product_id"`
	StoreID int    `json:"store_id,omitempty"`
	Name    string `json:"product_name"`
}

var ProductGraphqlObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Product",
	Fields:      graphql.BindFields(Product{}),
	Description: "Product object model",
})

func (p *Product) MakeFields() graphql.Fields {
	productField := make(graphql.Fields)

	productField["product_id"] = &graphql.Field{
		Name: "ProductID",
		Type: graphql.Int,
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			if product, ok := p.Source.(*Product); ok {
				i = product.ID
			} else {
				i = p.Source.(Product).ID
			}
			return
		},
	}

	productField["product_name"] = &graphql.Field{
		Name: "ProductName",
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			if product, ok := p.Source.(*Product); ok {
				i = product.Name
			} else {
				i = p.Source.(Product).Name
			}
			return
		},
	}

	return productField
}
