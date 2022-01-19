package persistence

import (
	"context"

	"github.com/etrnal70/kantin-backend/pkg/constants/model"
	"github.com/etrnal70/kantin-backend/pkg/storage"
	"github.com/gofrs/uuid"
)

func SellerRegister(data *model.SellerRegister, db storage.Database) (model.Seller, error) {
	row := db.DbPool.QueryRow(context.Background(),
		"INSERT INTO kantin.seller(name, email, password) VALUES($1, seller($2), crypt($3, gen_salt('bf', 8)), $4) RETURNING id, name, email,",
		data.Name, data.Email, data.Password)

	var seller model.Seller
	if err := row.Scan(&seller.ID, &seller.Name, &seller.Email); err != nil {
		return model.Seller{}, nil
	}

	return seller, nil
}

func SellerLogin(data *model.SellerLogin, db storage.Database) (*uuid.UUID, error) {
	row := db.DbPool.QueryRow(context.Background(),
		"SELECT email, password FROM kantin.seller WHERE email = $1 AND password = crypt($2, (SELECT password FROM kantin.seller WHERE email = $1))", data.Email,
	)

	var id uuid.UUID
	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	return &id, nil
}

func SellerGetAccount(seller_id uuid.UUID, db storage.Database) (model.Seller, error) {
	row := db.DbPool.QueryRow(context.Background(),
		"SELECT id, firstname, lastname, email FROM kantin.seller WHERE id= $1", seller_id,
	)

	var res model.Seller
	if err := row.Scan(&res); err != nil {
		return model.Seller{}, err
	}

	return res, nil
}

func SellerUpdateAccount(seller_id uuid.UUID, data *model.Seller, db storage.Database) (model.Seller, error) {
	row := db.DbPool.QueryRow(context.Background(),
		"UPDATE kantin.seller SET name = $1, email = $2, password = $3, logo = $4 WHERE id = $5 RETURNING id, name, email, password, logo", data.Name, data.Email, data.Password, data.Logo, seller_id,
	)
	var res model.Seller
	if err := row.Scan(&res); err != nil {
		return model.Seller{}, err
	}

	return res, nil
}

func GetCategory(db storage.Database) ([]model.Category, error) {
	rows, err := db.DbPool.Query(context.Background(),
		"SELECT * FROM kantin.category ORDER BY id",
	)
	if err != nil {
		return nil, err
	}

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.ID, &category.Value); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func GetSeller(db storage.Database) ([]model.SellerPublic, error) {
	rows, err := db.DbPool.Query(context.Background(),
		"SELECT id, name, logo FROM kantin.seller ORDER BY id",
	)
	if err != nil {
		return nil, err
	}

	var sellers []model.SellerPublic
	for rows.Next() {
		var seller model.SellerPublic
		if err := rows.Scan(&seller.ID, &seller.Name, &seller.Logo); err != nil {
			return nil, err
		}

		sellers = append(sellers, seller)
	}

	return sellers, nil
}

func GetSellerFood(seller_id uuid.UUID, db storage.Database) ([]model.MenuItem, error) {
	rows, err := db.DbPool.Query(context.Background(),
		"SELECT m.seller_id, m.id, m.name, m.description, m.category, m.price FROM kantin.seller s INNER JOIN kantin.menuitem m on $1 = m.seller_id", seller_id,
	)
	if err != nil {
		return nil, err
	}

	menuitems := []model.MenuItem{}

	for rows.Next() {
		var menu model.MenuItem
		rows.Scan(&menu.ID, &menu.Seller_ID, &menu.Name, &menu.Description, &menu.Category, &menu.Price)
		menuitems = append(menuitems, menu)
	}

	return menuitems, nil
}

// func GetFoodDetail(seller_id uuid.UUID, food_id uuid.UUID, db storage.Database) (model.MenuItem, error) {
// }

// func SellerAddFood(data []model.MenuItem, db storage.Database) (string, error) {}

// func SellerUpdateFood(data *model.MenuItem, db storage.Database) (model.MenuItem, error) {}

// func SellerGetFood(seller_id string, db storage.Database) ([]model.MenuItem, error) {}

// func SellerDeleteFood(seller_id string, data *[]model.MenuItem, db storage.Database) ([]model.MenuItem, error) {}

// func SellerGetOrder(seler_id string, db storage.Database) ([]model.Order, error) {}

// func SellerUpdateOrder(order_id string, db storage.Database) ([]model.Order, error) {}
