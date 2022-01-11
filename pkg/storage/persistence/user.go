package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/etrnal70/kantin-backend/pkg/constants/model"
	"github.com/etrnal70/kantin-backend/pkg/storage"
)

func UserRegister(data *model.UserRegister, db storage.Database) (int, error) {
	row := db.DbPool.QueryRow(context.Background(),
		"INSERT INTO kantin.user(firstname, lastname, email, password, createdAt) VALUES($1, $2, $3, $4, $5) RETURNING id", data.Firstname, data.Lastname, data.Email, data.Password, time.Now(),
	)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func UserLogin(data *model.UserLogin, db storage.Database) (*model.UserLogin, error) {
	rows, sqlerr := db.DbPool.Query(context.Background(),
		"SELECT email, password FROM kantin.user WHERE email = $1", data.Email,
	)
	if sqlerr != nil {
		return nil, fmt.Errorf("%w", sqlerr)
	}

	var res model.UserLogin
	for rows.Next() {
		rows.Scan(&res.Email, &res.Password)
	}

	return &res, nil
}

func UserGetAccount(user_id string, db storage.Database) (model.User, error) {
	row := db.DbPool.QueryRow(context.Background(),
		"SELECT id, firstname, lastname, email FROM kantin.user WHERE id = $1", user_id,
	)

	var res model.User
	if err := row.Scan(&res); err != nil {
		return model.User{}, err
	}

	return res, nil
}

func UserUpdateAccount(data *model.User, db storage.Database) (model.User, error) {
	row := db.DbPool.QueryRow(context.Background(),
		"UPDATE kantin.user SET firstname = $1, lastname = $2, email = $3, password = $4 WHERE id = $5 RETURNING id, firstname, lastname, email, password", data.Firstname, data.Lastname, data.Email, data.Password, data.ID,
	)

	var res model.User
	if err := row.Scan(&res); err != nil {
		return model.User{}, err
	}

	return res, nil
}

// func UserGetCategory(db storage.Database) ([]model.Category, error) {}

// func UserGetSeller(db storage.Database) ([]model.Seller, error) {}

// func UserGetSellerFood(seller_id string, db storage.Database) ([]model.MenuItem, error) {}

// func UserGetFoodDetail(food_id string, db storage.Database) (*model.MenuItem, error) {}

// func UserMakeOrder(data *model.Order, db storage.Database) (string, error) {}

// func UserGetOrderStatus(order_id string, db storage.Database) (*model.Status, error) {}

// func UserRequestOrderCancel(db storage.Database) (string, error) {}

// func UserGetOrders(user_id string, db storage.Database) ([]model.Order, error) {}
