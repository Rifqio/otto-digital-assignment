
# Backend Task Assignment

This project follows a structured approach to organizing Go code, making it easy to navigate and maintain. Below is a breakdown of the folder structure and its purpose.

## Folder Structure

```
├── database       # Contains database connection and migrations
├── dto            # Data Transfer Objects (DTOs) for request/response structures
├── handler        # HTTP handlers (controllers) for handling API requests
├── repositories   # Database interaction layer using GORM
├── routes         # API route definitions
├── service        # Business logic layer
├── utils          # Utility functions (helpers, error handling, etc.)
├── .env.example   # Environment variables file (Rename to .env)
├── go.mod         # Go module file
├── go.sum         # Go dependencies checksum
├── main.go        # Application entry point
├── README.md      # Project documentation
```

## Database Schema
Provided in the folder database and in the file schema.png

## Local Installation
Prerequisites

- Go installed on your system (version 1.24 recommended)

- A running MySQL database
    
    1. Install Dependencies
    ```
    go mod tidy
    ```
    2. Configure env for database url
    ```
    cp .env.example .env
    ```
    3. Install db migrate for migrations you can check out their documentation here https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md
    
    4. Run the migration command from backend project directory make sure your database is running
    ```
    migrate -path database/migrations -database "mysql://{username}:{password}@tcp(127.0.0.1:3306)/{dbname}" up
    ```

    5. Run the project by typing
    ```
    go run main.go
    ```
  
## Docker Installation
    1. Make sure docker engine is already running
    2. Run this command on the backend project directorys
    ```
    docker-compose up
    ```


## API Reference

#### Base Response
Base response for this project as follows

```
 {
     "success": boolean,
     "message": string,
     "data": <T>
 }
```

#### Create new brand

```http
  POST /api/v1/brand
```

| Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `brandName` | `string` | **Required** |

#### Create new voucher

```http
  POST /api/v1/voucher
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `code`      | `string` | **Required** |
| `point`      | `int` | **Required**.  |
| `validAt`      | `string` | **Required**.  |
| `expiredAt`      | `string` | **Required**.  |
| `brandId`      | `int` | **Required**.  |
| `createdAt`      | `string` | **Required**.  |
| `updatedAt`      | `string` | **Required**.  |


#### Get voucher details

```http
  GET /api/v1/voucher?id={id}
```

| Param | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required** |


#### Get voucher by brand id

```http
  GET /api/v1/voucher?id={id}
```

| Param | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required** |

#### Create redemption voucher

```http
  POST /api/v1/transaction/redemption
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `customerEmail`      | `string` | **Required** |
| `vouchers`      | `Array<{ voucherCode: string, quantity: int }>` | **Required** |

#### Get transaction history

```http
  POST /api/v1/transaction/redemption?transactionId={transactionId}
```

| Param | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `transactionId`      | `string` | **Required** |

