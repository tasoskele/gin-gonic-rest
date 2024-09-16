# Products RESTful CRUD API

## Description
- This is a simple CRUD API for products

- It is built with the Go-Gin framework and uses the RESTful API architecture.

- ...

## Setup instructions
- **Docker** .....

## Model Structure

### products
| Field       | Type         | Key     | Description |
|-------------|--------------|---------|-------------|
| id          | int          | Primary | Unique identifier for each product     |
| price       | float64      |         | The price of the product               |
| remaining   | int          |         | The remaining quantity of the product  |
| height      | float64      |         | The height of the product              |
| weight      | float64      |         | The weight of the product              |
| some_field  | -------      |         | -------------------------              |
| created     | datetime     |         | Created timestamp                      |
| updated     | datetime     |         | Updated timestamp                      |
| deleted     | datetime     |         | Deleted timestamp(?)                   |

## API endpoints
- products
  - GET /products
    - return all product entries
  - GET /products/{product_id}
    - return single product information
  - POST /products
    - create a new product entry
  - PUT /products/{product_id}
    - update information for a specific product
  - DELETE /products/{product_id}
    - delete a specific product
