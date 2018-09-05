# Demo GraphQL in Go

### Install dependencies
Using `glide` (https://github.com/Masterminds/glide)

```sh
$ glide install
```

### Build & Run

```sh
$ go build && ./demo-graphql
```

### Test

```sh
GET http://localhost:8080/graphql/store?query={store(store_id:5){store_id,store_name,products{product_id,product_name}}} 
```

Response:
```json
{
    "data": {
        "store": {
            "products": [
                {
                    "product_id": 1,
                    "product_name": "Hape seken"
                },
                {
                    "product_id": 2,
                    "product_name": "Hape baru"
                }
            ],
            "store_id": 5,
            "store_name": "Pantau Cocok Bayar"
        }
    }
}
```