package router

import (
	"net/http"

	handler "github.com/etrnal70/kantin-backend/pkg/handlers"
	middleware "github.com/etrnal70/kantin-backend/pkg/middlewares"
	"github.com/etrnal70/kantin-backend/pkg/storage"
	"github.com/gin-gonic/gin"
)

func UserRouter(db *storage.Database) http.Handler {
	r := gin.Default()
	r.Use(middleware.DbMiddleware(db))
	// user := r.Group("/user")
	r.POST("/user/register", handler.UserRegister)
	r.POST("/user/login", handler.UserLogin)
	r.GET("/user/:user_id", handler.UserGetAccount)
	r.PUT("/user/:user_id", handler.UserUpdateAccount)
	r.GET("/user/:user_id/order", handler.UserGetOrders)
	r.POST("/user/:user_id/order", handler.UserMakeOrder)
	r.GET("/user/:user_id/order/:order_id", handler.UserGetOrderStatus)
	r.POST("/user/:user_id/order/:order_id", handler.UserRequestOrderCancel)

	return r
}
