# Arithmetic Calculator Backend

Author: Ricardo Fabila

This project can be easily build using other technologies such as clojure, if requested 😁.

## Live Version

A live version of this project can be found at the following URL: [link-here](#). It is hosted on AWS.

**Credentials:**

- **User**: newuser@example.com
- **Password**: password123

> Note: I added a registration page as an extra feature, allowing new users to create accounts if needed.

## 🏃‍♂️ How to Run the Project Locally

To run the project locally, follow these steps:

1. Clone the project repository.
2. Run the command: `go run main.go`

**Requirements:**

- Go version 1.23 or greater.

## 💾 Database

This project uses SQLite as the database since it was the simplest solution for a quick task. If this were a real
production application, I would opt for PostgreSQL hosted on AWS RDS, and use a Docker instance for local development.

The project uses GORM (which is an ORM), which automatically creates the database file in the root of the project when
running
the application. For more control and less "magic," I would recommend using a library like `sqlx` along with a migration
tool like Flyway, which is the setup we use at my current company to manage millions of records per customer.

## 🧪 Testing

This project includes automated tests for each API endpoint to ensure functionality and reliability.

Run the tests with:

´go test ./... -count=1´

* ´-count=1´ is used to prevent cached tests

## API Examples

Below are a few cURL commands to try out the API while developing locally. You can also use them against the live
version by replacing the base URL.

### Register Endpoint (POST /register)

```sh
curl -X POST "http://localhost:8080/register" \
-H "Content-Type: application/json" \
-d '{
  "username": "newuser@example.com",
  "password": "password123"
}'
```

### Login Endpoint (POST /login)

```sh
curl -X POST "http://localhost:8080/login" \
-H "Content-Type: application/json" \
-d '{
  "username": "newuser@example.com",
  "password": "password123"
}'
```

### Perform Operations (POST /api/v1/operation)

#### Addition Operation

```sh
curl -X POST "http://localhost:8080/api/v1/operation" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <token>" \
-d '{
  "operation": "addition",
  "number1": 5,
  "number2": 3
}'
```

#### Subtraction Operation

```sh
curl -X POST "http://localhost:8080/api/v1/operation" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <token>" \
-d '{
  "operation": "subtraction",
  "number1": 10,
  "number2": 4
}'
```

#### Square Root Operation

```sh
curl -X POST "http://localhost:8080/api/v1/operation" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <token>" \
-d '{
  "operation": "square_root",
  "number1": 16
}'
```

#### Random String Operation

```sh
curl -X POST "http://localhost:8080/api/v1/operation" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <token>" \
-d '{
  "operation": "random_string",
  "length": 15
}'
```

### Get Records (GET /api/v1/records)

```sh
curl -X GET "http://localhost:8080/api/v1/records?page=1&limit=10" \
-H "Authorization: Bearer <token>"
```

### Delete a Record (DELETE /api/v1/records/:id)

```sh
curl -X DELETE "http://localhost:8080/api/v1/records/1" \
-H "Authorization: Bearer <token>"
```

> Replace `<token>` with a valid JWT token obtained from the login endpoint.


