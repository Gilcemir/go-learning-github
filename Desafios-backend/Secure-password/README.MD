# Secure Password API

This project provides a simple API to validate passwords. It exposes one route that checks the validity of the password provided in the request body. Depending on the input, it returns appropriate HTTP status codes.

## Features

- **POST /secure-password**: An endpoint to validate passwords.
- Returns:
  - `204 No Content`: If the password is valid.
  - `400 Bad Request`: For invalid requests such as:
    - Invalid password format.
    - Missing or incorrect body in the request.
    - Usage of a method other than `POST`.

## Requirements

- **Go**: Version 1.23 or higher.

## Project Structure

The main components of the project are:

- `cmd/`: Contains the main application entry point.
- `internal/validator/`: Contains the password validation logic.

## How to Run

1. Clone the repository:

   ```bash
   git clone https://github.com/Gilcemir/go-learning-github.git
   cd go-learning-github/Desafios-backend/Secure-password
   ```
2. Ensure Go modules are enabled and install any dependencies:
    ```bash
    go mod tidy
    ```
3. Run the application
   ```bash
   make serve
   ```
   or
   ```bash
   go run cmd/main.go
   ```

# API Documentation

## Endpoint: `POST /secure-password`

### Description
This endpoint accepts a password in the request body and validates it.

### Content-Type
`application/json`

You can use the test.http as example

### Request Body
The body should be a JSON object with the following structure:

```json
{
  "password": "your-password-here"
}
```

# Running tests

   ```bash
   make test
   ```
   or
   ```bash
   go test ./...
   ```

