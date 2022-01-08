package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/etrnal70/kantin-backend/pkg/constants/model"
	"github.com/etrnal70/kantin-backend/pkg/storage"
)

// Seller can:
// 1. Register
func SellerRegister(data *model.SellerRegister, db storage.Database) (string, error) {
	row, sqlerr := db.DbPool.Query(context.Background(),
		"INSERT INTO kantin.seller(name, email, password, createdAt) VALUES($1, $2, $3, $4) RETURNING id",
		data.Name, data.Email, data.Password, time.Now())
	if sqlerr != nil {
		return "", fmt.Errorf("%w", sqlerr)
	}

	var id string
	row.Scan(&id)

	return id, nil
}

// 2. Login
func SellerLogin(data *model.SellerLogin, db storage.Database) (model.SellerLogin, error) {
	row, sqlerr := db.DbPool.Query(context.Background(),
		"SELECT email, password FROM kantin.seller WHERE email = $1", data.Email,
	)
	if sqlerr != nil {
		return model.SellerLogin{}, fmt.Errorf("%w", sqlerr)
	}

	var res model.SellerLogin
	row.Scan(&res)

	return res, nil
}

// 3. Get account details
func SellerGetAccount(seller_id string, db storage.Database) (model.Seller, error) {
	row, sqlerr := db.DbPool.Query(context.Background(),
		"SELECT id, firstname, lastname, email FROM kantin.seller WHERE id= $1", seller_id,
	)
	if sqlerr != nil {
		return model.Seller{}, fmt.Errorf("%w", sqlerr)
	}

	var res model.Seller
	row.Scan(&res)

	return res, nil
}

// 4. Change account details
func SellerUpdateAccount(data *model.Seller, db storage.Database) (model.Seller, error) {
  row, sqlerr := db.DbPool.Query(context.Background(),
    "UPDATE kantin.seller SET name = $1, email = $2, password = $3, logo = $4 WHERE id = $5 RETURNING id, name, email, password, logo", data.Name, data.Email, data.Password, data.Logo,
  )
  if sqlerr != nil {
    return model.Seller{}, fmt.Errorf("%w", sqlerr)
  }

  var res model.Seller
  row.Scan(&res)

  return res, nil
}

// 5. Add food
// func SellerAddFood(data []model.MenuItem, db storage.Database) (string, error) {}

// 6. Update food
// func SellerUpdateFood(data *model.MenuItem, db storage.Database) (model.MenuItem, error) {}

// 7. List food
// func SellerGetFood(seller_id string, db storage.Database) ([]model.MenuItem, error) {}

// 8. Delete food
// func SellerDeleteFood(seller_id string, data *[]model.MenuItem, db storage.Database) ([]model.MenuItem, error) {}

// 9. List all orders
// func SellerGetOrder(seler_id string, db storage.Database) ([]model.Order, error) {}

// 10. Change order status
// func SellerUpdateOrder(order_id string, db storage.Database) ([]model.Order, error) {}
