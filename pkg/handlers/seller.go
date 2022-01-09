package handler

import (
	"fmt"
	"net/http"

	"github.com/etrnal70/kantin-backend/pkg/constants/model"
	"github.com/etrnal70/kantin-backend/pkg/storage"
	"github.com/etrnal70/kantin-backend/pkg/storage/persistence"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SellerRegister(c *gin.Context) {
	db, status := c.MustGet("db").(storage.Database)
	if !status {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": status,
		})
	}

	var data model.SellerRegister
	c.BindJSON(&data)

	// Hash password before storing
	hashedpass, hashederr := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if hashederr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Unable to register seller : %s", hashederr),
		})
	}

	data.Password = string(hashedpass)
	res, err := persistence.SellerRegister(&data, db)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	// TODO: Should return token, stil missing middleware
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"token":   res,
	})
}

func SellerLogin(c *gin.Context) {
	db, status := c.MustGet("db").(storage.Database)
	if !status {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": status,
		})
	}

	var data model.SellerLogin
	c.BindJSON(&data)

	dbData, err := persistence.SellerLogin(&data, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbData.Password), []byte(data.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Password mismatch",
		})
		return
	}

	// TODO: Should return token, still missing middleware
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"token":   dbData,
	})
}

func SellerGetAccount(c *gin.Context) {
	db, status := c.MustGet("db").(storage.Database)
	if !status {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": status,
		})
	}

	seller_id := c.Param("seller_id")
	res, err := persistence.SellerGetAccount(seller_id, db)
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

func SellerUpdateAccount(c *gin.Context) {
	db, status := c.MustGet("db").(storage.Database)
	if !status {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": status,
		})
	}

	var data model.Seller
	c.BindJSON(&data)

	res, err := persistence.SellerUpdateAccount(&data, db)
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

// func SellerAddFood(c *gin.Context) {
// 	db, status := c.MustGet("db").(storage.Database)
// 	if !status {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": status,
// 		})
// 	}
// }

// func SellerUpdateFood(c *gin.Context) {
// 	db, status := c.MustGet("db").(storage.Database)
// 	if !status {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": status,
// 		})
// 	}
// }

// func SellerGetFoods(c *gin.Context) {
// 	db, status := c.MustGet("db").(storage.Database)
// 	if !status {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": status,
// 		})
// 	}
// }

// func SellerDeleteFood(c *gin.Context) {
// 	db, status := c.MustGet("db").(storage.Database)
// 	if !status {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": status,
// 		})
// 	}
// }

// func SellerGetOrders(c *gin.Context) {
// 	db, status := c.MustGet("db").(storage.Database)
// 	if !status {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": status,
// 		})
// 	}
// }

// func SellerUpdateOrder(c *gin.Context) {
// 	db, status := c.MustGet("db").(storage.Database)
// 	if !status {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": status,
// 		})
// 	}
// }
