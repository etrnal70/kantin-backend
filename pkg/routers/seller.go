package router

import (
	"net/http"

	handler "github.com/etrnal70/kantin-backend/pkg/handlers"
	middleware "github.com/etrnal70/kantin-backend/pkg/middlewares"
	"github.com/etrnal70/kantin-backend/pkg/storage"
	"github.com/gin-gonic/gin"
)

func SellerRouter(db *storage.Database) http.Handler {
	r := gin.Default()
	r.Use(middleware.DbMiddleware(db))
	// seller := r.Group("/seller")
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

	return r
}
