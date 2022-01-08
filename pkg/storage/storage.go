package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type db_connection interface {
	Close()
}

type Database struct {
	DbPool *pgxpool.Pool
}

// TODO: Generalize usage to be able to use redis
func InitDb() (*Database, error) {
	dbname := os.Getenv("DB_NAME")
	// host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

  db_url := fmt.Sprintf("host=db database=%s user=%s password=%s sslmode=disable", dbname, username, password)
  fmt.Println(db_url)

	dbPool, err := pgxpool.Connect(context.Background(), db_url)
	if err != nil {
		fmt.Printf("Unable to connect to database : %s\n", err)
		return nil, err
	}

	var res Database
	res.DbPool = dbPool
	return &res, nil
}

func (dbConn Database) Close() {
	dbConn.DbPool.Close()
}
