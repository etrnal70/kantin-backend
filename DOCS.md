# Documentation

## API

### User

```go
POST("/user/register", handler.UserRegister)
POST("/user/login", handler.UserLogin)
GET("/user/:user_id", handler.UserGetAccount)
PUT("/user/:user_id", handler.UserUpdateAccount)
GET("/user/:user_id/order", handler.UserGetOrders)
POST("/user/:user_id/order", handler.UserMakeOrder)
GET("/user/:user_id/order/:order_id", handler.UserGetOrderStatus)
POST("/user/:user_id/order/:order_id", handler.UserRequestOrderCancel)
```

### Seller / General

```go
r.GET("/seller")
r.GET("/seller/menu")
r.GET("/seller/menu/:id")
r.GET("/category")

r.POST("/seller/register", handler.SellerRegister)
r.POST("/seller/login", handler.SellerLogin)
r.GET("/seller/:seller_id", handler.SellerGetAccount)
r.PUT("/seller/:seller_id", handler.SellerUpdateAccount)
// r.POST("/seller/food", handler.SellerAddFood)
// r.POST("/seller/food/:food_id", handler.SellerUpdateFood)
// r.GET("/seller/food", handler.SellerGetFoods)
// r.DELETE("/seller/food/:food_id", handler.SellerDeleteFood)
// r.GET("/seller/order", handler.SellerGetOrders)
// r.PUT("/seller/order/:order_id", handler.SellerUpdateOrder)
```

## Table Details

### Category

| id  | category |
| --- | -------- |
| 1   | FOOD     |
| 2   | SNACK    |
| 3   | BEVERAGE |
| 4   | DESSERT  |

### Status

TODO: Discuss this

| id  | status                   |
| --- | ------------------------ |
| 1   | WAITING_APPROVAL         |
| 2   | PROCESSING               |
| 3   | REQUEST_FOR_CANCELLATION |
| 4   | READY                    |
| 5   | DONE                     |
| 6   | CANCELLED                |
