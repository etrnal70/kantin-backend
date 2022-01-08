package middleware

import (
	"github.com/etrnal70/kantin-backend/pkg/storage"
	"github.com/gin-gonic/gin"
)

func DbMiddleware(db *storage.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", &db)
		c.Next()
	}
}
