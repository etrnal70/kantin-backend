package router

import (
	"net/http"

	handler "github.com/etrnal70/kantin-backend/pkg/handlers"
	middleware "github.com/etrnal70/kantin-backend/pkg/middlewares"
	"github.com/etrnal70/kantin-backend/pkg/storage"
	"github.com/gin-gonic/gin"
)

func UserRouter(db storage.Database) http.Handler {
	r := gin.Default()
  userHandler := new(handler.UserHandler)
  userHandler.Conn = db
	r.POST("/user/register", userHandler.UserRegister)

  protectedUser := r.Group("/user", middleware.AuthMiddleware())
	protectedUser.POST("/user/login", userHandler.UserLogin)
	protectedUser.GET("/user/:user_id", userHandler.UserGetAccount)
	protectedUser.PUT("/user/:user_id", userHandler.UserUpdateAccount)
	protectedUser.GET("/user/:user_id/order", userHandler.UserGetOrders)
	protectedUser.POST("/user/:user_id/order", userHandler.UserMakeOrder)
	protectedUser.GET("/user/:user_id/order/:order_id", userHandler.UserGetOrderStatus)
	protectedUser.POST("/user/:user_id/order/:order_id", userHandler.UserRequestOrderCancel)

	return r
}
