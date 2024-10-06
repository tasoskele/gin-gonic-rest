# Products RESTful CRUD API

## Description
- This is a simple CRUD API for products

- It is built with the Go-Gin framework and uses the RESTful API architecture.

- The application is dockerized using Docker Compose for easy deployment.

- It uses pagination with params like so:
| Key   | Value |
| page  |  2    |
| limit |  5    |

## Setup instructions
**Requirements**
- Docker
- Docker Compose

**Installation**
1. Clone the repository and change directory into it
2. run: ```docker-compose up --build```

## Model Structure

### products
| Field       | Type         | Key     | Description |
|-------------|--------------|---------|-------------|
| id          | int          | Primary | Unique identifier for each product     |
| title       | string       |         | The title of the product               |
| body        | string       |         | The description of the product         |
| price       | float64      |         | The price of the product               |
| created     | datetime     |         | Created timestamp                      |
| updated     | datetime     |         | Updated timestamp                      |
| deleted     | datetime     |         | Deleted timestamp                      |

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
