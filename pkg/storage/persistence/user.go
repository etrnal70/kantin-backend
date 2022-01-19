package persistence

import (
	"context"
	"fmt"

	"github.com/etrnal70/kantin-backend/pkg/constants/model"
	"github.com/etrnal70/kantin-backend/pkg/storage"
	"github.com/gofrs/uuid"
)

func UserRegister(data *model.UserRegister, db storage.Database) (model.User, error) {
	row := db.DbPool.QueryRow(context.Background(),
		"INSERT INTO kantin.user(firstname, lastname, email, password) VALUES($1, $2, lower($3), crypt($4, gen_salt('bf', 8))) RETURNING id, firstname, lastname, email", data.Firstname, data.Lastname, data.Email, data.Password,
	)

	var user model.User
	if err := row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email); err != nil {
		return model.User{}, err
	}

	return user, nil
}

func UserLogin(data *model.UserLogin, db storage.Database) (*uuid.UUID, error) {
	row, sqlerr := db.DbPool.Query(context.Background(),
		"SELECT id FROM kantin.user WHERE email = $1 AND password = crypt($2, (SELECT password FROM kantin.user WHERE email = $1))", data.Email, data.Password,
	)
	if sqlerr != nil {
		return nil, fmt.Errorf("%w", sqlerr)
	}

	var id uuid.UUID
	if err := row.Scan(id); err != nil {
    return nil, err
  }

	return &id, nil
}

func UserGetAccount(user_id uuid.UUID, db storage.Database) (model.User, error) {
	row := db.DbPool.QueryRow(context.Background(),
		"SELECT id, firstname, lastname, email FROM kantin.user WHERE id = $1", user_id,
	)

	var res model.User
	if err := row.Scan(&res.ID, &res.Firstname, &res.Lastname, &res.Email); err != nil {
		return model.User{}, err
	}

	return res, nil
}

func UserUpdateAccount(user_id uuid.UUID, data *model.User, db storage.Database) (model.User, error) {
	row := db.DbPool.QueryRow(context.Background(),
		"UPDATE kantin.user SET firstname = $1, lastname = $2, email = $3, password = $4 WHERE id = $5 RETURNING id, firstname, lastname, email, password", data.Firstname, data.Lastname, data.Email, data.Password, user_id,
	)

	var res model.User
	if err := row.Scan(&res.ID, &res.Firstname, &res.Lastname, &res.Email, &res.Password); err != nil {
		fmt.Printf("Error: %s", err)
		return model.User{}, err
	}

	return res, nil
}

// func UserMakeOrder(data *model.Order, db storage.Database) (string, error) {}

// func UserGetOrderStatus(order_id string, db storage.Database) (*model.Status, error) {}

// func UserRequestOrderCancel(db storage.Database) (string, error) {}

// func UserGetOrders(user_id string, db storage.Database) ([]model.Order, error) {}
