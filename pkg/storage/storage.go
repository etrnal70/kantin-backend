package storage

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type db_connection interface {
	Close()
}

type Database struct {
	DbPool *pgxpool.Pool
}

// TODO: Generalize usage to be able to use redis
func InitDb(timeout time.Duration) (*Database, error) {
	var res Database
	dbname := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	db_url := fmt.Sprintf("host=db database=%s user=%s password=%s sslmode=disable", dbname, username, password)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeoutReached := time.After(timeout)
	for {
		select {
		case <-timeoutReached:
			return nil, fmt.Errorf("Database timeout after %s", timeout)

		case <-ticker.C:
			dbPool, err := pgxpool.Connect(context.Background(), db_url)
			if err == nil {
				res.DbPool = dbPool
				return &res, nil
			}
			fmt.Printf("Waiting for database connection...\n")
		}
	}
}

func (dbConn Database) Close() {
	dbConn.DbPool.Close()
}
