
# Email Validator Service
## What is this repository for? ##
This is email validator service using [emailvalidator](https://github.com/sufimalek/emailvalidator) library

## Project Layout ##

The repository uses the following folder structure:

```
email-validator-service/
|-- cmd/
|   |-- server/
|       |-- main.go
|   |-- server/
|       |-- handlers.go
|       |-- middleware.go
|       |-- routes.go
|-- config/
|   |-- config.go
|-- go.mod
|-- go.sum
|-- README.md
```

## Library used for this service

### [EmailValidator] https://github.com/sufimalek/emailvalidator


## How do I get set up? ##

#### Clone project

git clone email-validator-service


## How to run service ##

```sh
cd email-validator-service
go run cmd/server/main.go
```

and this service will be available on 8080 port

## Running the Service using Make

1. Install Dependencies:
    ```sh
    make deps
    ```
2. Build the Service:
    ```sh
    make build
    ```
3. Run the Service:
    ```sh
    make run
    ```
4. Test the Service:
    ```sh
    make test
    ```
5. Clean the Build:
    ```sh
    make clean
    ```


## How to call service for API ##

```curl
curl --location 'localhost:8080/api/validate' \
--header 'Content-Type: application/json' \
--data-raw '{"email":"abc@mail.com"}'
```

## API Request

1. Endpoint: /api/validate
2. Method: POST
3. Request Body:

```json
{
  "email": "example@example.com"
}
```

4. Response:
```json
{
  "is_valid": true,
}
```
or
```json
{
  "is_valid": false,
  "message": "Email validation failed"
}
```
