# Payment service

Change config values in app.env to your configuration.
```
STRIPE_KEY=***
```

## Endpoints

```
GET  localhost:8080/v1/payments/:id
GET  localhost:8080/v1/payments?offset=0&limit=10
POST localhost:8080/v1/payments
```

Payment JSON:
```
{
  "booking_id": 1,
  "price": 250.1,
  "paid": false
}
```

## gRPC API

Payment PROTO:

```
{
    "payment": {
        "booking_id": 1,
        "paid": true,
        "price": 512.5
    }
}
```

### Modifying Payment Proto

If `proto/payment.proto` is modified, run:

```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto
```

If protoc is not found by bash, run:

```
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
```