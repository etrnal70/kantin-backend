package handler

import (
	"fmt"
	"net/http"

	"github.com/etrnal70/kantin-backend/pkg/constants/model"
	middleware "github.com/etrnal70/kantin-backend/pkg/middlewares"
	"github.com/etrnal70/kantin-backend/pkg/storage"
	"github.com/etrnal70/kantin-backend/pkg/storage/persistence"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type SellerHandler struct {
	Conn storage.Database
}

func (db *SellerHandler) SellerRegister(c *gin.Context) {
	var data model.SellerRegister
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	res, err := persistence.SellerRegister(&data, db.Conn)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	token, err := middleware.GenerateToken(res.ID.String())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"token":   res,
	})
}

func (db *SellerHandler) SellerLogin(c *gin.Context) {
	var data model.SellerLogin
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	id, err := persistence.SellerLogin(&data, db.Conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	token, err := middleware.GenerateToken(id.String())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func (db *SellerHandler) SellerGetAccount(c *gin.Context) {
	seller_id, _ := uuid.FromString(c.Param("user_id"))
	res, err := persistence.SellerGetAccount(seller_id, db.Conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    res.ID,
		"email": res.Email,
		"name":  res.Name,
		"logo":  res.Logo,
	})
}

func (db *SellerHandler) SellerUpdateAccount(c *gin.Context) {
	seller_id, _ := uuid.FromString(c.Param("user_id"))
	var data model.Seller
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	res, err := persistence.SellerUpdateAccount(seller_id, &data, db.Conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    res.ID,
		"email": res.Email,
		"name":  res.Name,
		"logo":  res.Logo,
	})
}

func (db *SellerHandler) GetCategory(c *gin.Context) {
	res, err := persistence.GetCategory(db.Conn)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"category": res,
	})
} // Marked
func (db *SellerHandler) GetSeller(c *gin.Context) {
	res, err := persistence.GetSeller(db.Conn)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"seller": res,
	})
} // Marked
func (db *SellerHandler) GetSellerFood(c *gin.Context) {
	seller_id, _ := uuid.FromString(c.Param("seller_id"))
	res, err := persistence.GetSellerFood(seller_id, db.Conn)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"seller_id": seller_id,
		"foods":     res,
	})
} // Marked
// func (db *SellerHandler) GetFoodDetail(c *gin.Context) {
// 	seller_id, _ := uuid.FromString(c.Param("seller_id"))
// 	food_id, _ := uuid.FromString(c.Param("food_id"))
// 	res, err := persistence.GetFoodDetail(seller_id, food_id, db.Conn)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": err,
// 		})
//
// 		return
// 	}
//
// 	c.JSON(http.StatusOK, gin.H{
// 		"menuitem": res,
// 	})
// } // Marked

func (db *SellerHandler) SellerAddFood(c *gin.Context) {}

func (db *SellerHandler) SellerUpdateFood(c *gin.Context) {}

func (db *SellerHandler) SellerGetFoods(c *gin.Context) {}

func (db *SellerHandler) SellerDeleteFood(c *gin.Context) {}

func (db *SellerHandler) SellerGetOrders(c *gin.Context) {}

func (db *SellerHandler) SellerUpdateOrder(c *gin.Context) {}
