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

### Parameters

#### Add new tax object

| No | Field | Type | Description |
| -- | ----- | ---- | ----------- |
| 1 | `name` | `string` | Tax object's name |
| 2 | `tax_code` | `string` | Tax object's category code, refer to [here](#tax-category-code) |
| 3 | `price` | `string` | Tax object's price amount |

#### Get tax objects summary

n/a

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
