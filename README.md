# :dollar: Tax Calculator

API for calculating your tax objects.

Supported tax category:

- Food & Beverage
- Entertainment
- Tobacco

## Run

Build normally:

```
$ go build -mod=vendor && ./tax-calc
```

Or run directly:

```
$ go run -mod=vendor main.go
```

Or use `docker`:

```
$ docker-compose up
```

## Routes

| No | Path | Method | Type | Description |
| -- | ---- | ------ | ---- | ----------- |
| 1 | `/bills` | `POST` | `x-www-form-urlencoded` | [Add new tax object](#add-new-tax-object) |
| 2 | `/bills` | `GET` | n/a | [Get tax objects summary](#get-tax-objects-summary) |

### Response

Response is always in such form:

```JSON
{
    "data": { .. },
    "error": " .. "
}
```

| No | Field | Type | Description |
| -- | ----- | ---- | ----------- |
| 1 | `data` | `object` | Response object |
| 2 | `error` | `string` | Error message if any |

Returned HTTP status code also depends on whether the request is successful:

- `200` - `OK`, success.
- `400` - `BAD_REQUEST`, usually invalid param.
- `500` - `INTERNAL_ERROR`, usually something wrong with the server.

### Add new tax object

#### Parameters

| No | Field | Type | Description |
| -- | ----- | ---- | ----------- |
| 1 | `name` | `string` | Tax object's name |
| 2 | `tax_code` | `string` | Tax object's category code, refer to [here](#tax-category-code) |
| 3 | `price` | `string` | Tax object's price amount |

#### Response

On success, response http status code 200 with generic JSON response is returned (can be ignored).

Sample response:

```JSON
{
    "data": null
}
```

### Get tax objects summary

#### Parameters

n/a

#### Response

Field explanation within `data`:

| No | Field | Type | Description |
| -- | ----- | ---- | ----------- |
| 1 | `bills` | `array` of `BillDetail` | List of all tax object's details |
| 2 | `price_subtotal` | `string` | Tax object's price subtotal |
| 3 | `tax_subtotal` | `string` | Tax object's tax subtotal |
| 4 | `grand_total` | `string` | Tax object's price and tax grand total |

Field explanation for `BillDetail`:

| No | Field | Type | Description |
| -- | ----- | ---- | ----------- |
| 1 | `name` | `string` | Tax object's name |
| 2 | `tax_code` | `string` | Tax object's category code |
| 3 | `type` | `string` | Tax object's category name |
| 4 | `refundable` | `boolean` | Tax object's  refundable indicator |
| 4 | `price` | `string` | Tax object's price amount |
| 4 | `tax` | `string` | Tax object's tax amount |
| 4 | `amount` | `string` | Tax object's total amount |

Sample response:

```json
{
    "data": {
        "bills": [
            {
                "name": "Big Mac",
                "tax_code": "2",
                "type": "Tobacco",
                "refundable": false,
                "price": "1000.00",
                "tax": "30.00",
                "amount": "1030.00"
            }
        ],
        "price_subtotal": "1000.00",
        "tax_subtotal": "30.00",
        "grand_total": "1030.00"
    }
}
```

### Tax Category Code

| No | Category | Code |
| -- | -------- | ---- |
| 1 | Food & Beverage | `1` |
| 2 | Entertainment | `2` |
| 3 | Tobacco | `3` |

## Sample request

Postman Collection: [https://www.getpostman.com/collections/28a52c1830b0d19113ed](https://www.getpostman.com/collections/28a52c1830b0d19113ed)

## Database structure

Refer to the `migration` directory [here](https://github.com/rudbast/tax-calc/tree/master/files/data/tax-calc/migration).

Field explanation is provided in `.sql` file in comments form.
