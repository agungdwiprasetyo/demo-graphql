# Demo GraphQL in Go

## Install dependencies
Using `glide` (https://github.com/Masterminds/glide)

```sh
$ glide install
```

## Build & Run

```sh
$ go build && ./demo-graphql
```

## Test

```sh
GET http://localhost:8080/store/graphql?query={store(store_id:5){store_id,store_name,products{product_id,product_name}}} 
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

## GraphQL Usage

Here is how you can use the queries:

<table>
    <tr>
        <th>Action</th>
        <th>Query</th>
    </tr>
    <tr>
        <td>
Get All Store:
        </td>
        <td>
            <pre>
{
    get_all_stores{
        store_id,
        store_name,
        products{
            product_id,
            product_name
        }
    }
}
            </pre>
        </td>
    </tr>
    <tr>
        <td>
Get Detail Store:
        </td>
        <td>
            <pre>
{
    get_store(store_id:1){
        store_id,
        store_name,
        products{
            product_id,
            product_name
        }
    }
}
            </pre>
        </td>
    </tr>
    <tr>
        <td>
Get All Product:
        </td>
        <td>
            <pre>
{
    get_all_product{
        product_id,
        product_name,
        store{
            store_id,
            store_name
        }
    }
}
            </pre>
        </td>
    </tr>
    <tr>
        <td>
Get Detail Product:
        </td>
        <td>
            <pre>
{
    get_product(product_id:2){
        product_id,
        product_name,
        store{
            store_id,
            store_name
        }
    }
}
            </pre>
        </td>
    </tr>
</table>