# Go Payment API

PaymentAPI is a simple API built using Go that allows users to submit, retrieve, update and delete payment information.

## Installation

- Clone the repository
- Run `go mod tidy` to install dependencies
- Set up environment variables in a .env file
- Run go run main.go to start the serve

## Dependencies

- gin-gonic/gin
- gorm.io/gorm
- github.com/joho/godotenv

## Endpoints

- POST `/payment` -  creates a new payment record in the database.
- PUT `/payment/:id` - updates a payment record in the database.
- GET `/allpayments` - returns a list of all payment records in the database.
- GET `/payment/:id` - returns a single payment record based on the provided ID.
- DELETE `/payment/:id` - deletes a single payment record based on the provided ID.

## Sample Responses:

- GET `/allpayments`:

```JSON
  {
    "data": [
        {
            "ID": 1,
            "CreatedAt": "2022-03-26T17:03:00.000+0000",
            "UpdatedAt": "2022-03-26T17:03:00.000+0000",
            "DeletedAt": null,
            "Amount": 100.00,
            "Quantity": 1.0,
            "Description": "Payment for item A",
            "ItemCode": "A123"
        },
        {
            "ID": 2,
            "CreatedAt": "2022-03-26T17:03:00.000+0000",
            "UpdatedAt": "2022-03-26T17:03:00.000+0000",
            "DeletedAt": null,
            "Amount": 50.00,
            "Quantity": 2.0,
            "Description": "Payment for item B",
            "ItemCode": "B456"
        }
    ]
}


```


- GET `/payment/:id`:

```JSON
{
    "ID": 1,
    "CreatedAt": "2022-03-26T17:03:00.000+0000",
    "UpdatedAt": "2022-03-26T17:03:00.000+0000",
    "DeletedAt": null,
    "Amount": 100.00,
    "Quantity": 1.0,
    "Description": "Payment for item A",
    "ItemCode": "A123"
}
```

  
