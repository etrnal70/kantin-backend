package router

import (
	"net/http"

	handler "github.com/etrnal70/kantin-backend/pkg/handlers"
	middleware "github.com/etrnal70/kantin-backend/pkg/middlewares"
	"github.com/etrnal70/kantin-backend/pkg/storage"
	"github.com/gin-gonic/gin"
)

func SellerRouter(db storage.Database) http.Handler {
	r := gin.Default()
	sellerHandler := new(handler.SellerHandler)
	sellerHandler.Conn = db

	r.GET("/category", sellerHandler.GetCategory)
	r.GET("/seller", sellerHandler.GetSeller)
	r.GET("/seller/:seller_id", sellerHandler.GetSellerFood)
	// r.GET("/seller/:seller_id/:menu_id", sellerHandler.GetFoodDetail)
	r.POST("/seller/register", sellerHandler.SellerRegister)

	protectedSeller := r.Group("/seller", middleware.AuthMiddleware())
	protectedSeller.POST("/seller/login", sellerHandler.SellerLogin)
	protectedSeller.GET("/seller/:seller_id", sellerHandler.SellerGetAccount)
	protectedSeller.PUT("/seller/:seller_id", sellerHandler.SellerUpdateAccount)
	protectedSeller.POST("/seller/food", sellerHandler.SellerAddFood)
	protectedSeller.POST("/seller/food/:food_id", sellerHandler.SellerUpdateFood)
	protectedSeller.GET("/seller/food", sellerHandler.SellerGetFoods)
	protectedSeller.DELETE("/seller/food/:food_id", sellerHandler.SellerDeleteFood)
	protectedSeller.GET("/seller/order", sellerHandler.SellerGetOrders)
	protectedSeller.PUT("/seller/order/:order_id", sellerHandler.SellerUpdateOrder)

	return r
}
