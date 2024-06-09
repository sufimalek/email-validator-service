# Email Validator Service

## Overview

The Email Validator Service is a comprehensive and production-ready service built in Go, designed to validate email addresses through various checks. It supports multiple validation mechanisms including syntax validation, role-based email checks, MX records verification, SMTP validation, and more.

## Features

- **Syntax Validation**: Ensures the email address follows standard email syntax rules.
- **Role-Based Email Check**: Identifies and rejects role-based email addresses (e.g., admin@example.com).
- **MX Record Verification**: Checks if the email's domain has valid MX records.
- **SMTP Validation**: Verifies the existence of the email address via SMTP.
- **Ban Words in Username Check**: Blocks email addresses with specific banned words in the username.
- **Blacklisted Emails Check**: Rejects email addresses that are blacklisted.
- **DNS Check**: Validates the domain's DNS records.
- **RFC Compliance Check**: Ensures the email address complies with RFC standards.
- **No RFC Warnings Check**: Ensures there are no RFC warnings.
- **Gravatar Check**: Checks if the email address is associated with a Gravatar.

## Installation

### Prerequisites

- Go 1.16+
- Docker (optional, for containerized deployment)

### Clone the Repository

```sh
git clone https://github.com/sufimalek/email-validator-service.git
cd email-validator-service
```

### Install Dependencies

```sh
make deps
```

### Configuration

Create a `.env` file in the root directory with the following content:

```
PORT=8080
```

### Build and Run

#### Using Makefile

```sh
make build
make run
```

#### Using Docker

Build the Docker image:

```sh
docker build -t email-validator-service .
```

Run the Docker container:

```sh
docker run -p 8080:8080 --env-file .env email-validator-service
```

## Usage

### API Endpoints

#### Validate Email

- **URL**: `/api/validate`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "email": "example@example.com"
  }
  ```
- **Response** (Valid Email):
  ```json
  {
    "is_valid": true,
  }
  ```
- **Response** (Invalid Email):
  ```json
  {
    "is_valid": false,
    "message": "Email validation failed"
  }
  ```

### Sample Request

You can use `curl` to make a request to the API:

```sh
curl -X POST http://localhost:8080/api/validate -H "Content-Type: application/json" -d '{"email": "test@example.com"}'
```

### Response

```json
{
  "is_valid": true,
}
```

## Project Structure

```
email-validator-service/
|-- cmd/
|   |-- server/
|       |-- main.go
|-- internal/
|   |-- server/
|       |-- handlers.go
|       |-- middleware.go
|       |-- routes.go
|-- config/
|   |-- config.go
|-- vendor/
|-- go.mod
|-- go.sum
|-- README.md
|-- Makefile
```

## Running Tests

To run tests, use the following command:

```sh
make test
```

## Contributing

We welcome contributions! Please fork the repository and submit a pull request.

### Steps to Contribute

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes and push to your fork.
4. Submit a pull request describing your changes.

## License

This project is licensed under the MIT License.

## Contact

For any inquiries or support, please open an issue on the [GitHub repository](https://github.com/sufimalek/email-validator-service) or contact us at suf_sap@hotmail.com.

---