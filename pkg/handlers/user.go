package handler

import (
	"fmt"
	"net/http"

	"github.com/etrnal70/kantin-backend/pkg/constants/model"
	"github.com/etrnal70/kantin-backend/pkg/storage"
	"github.com/etrnal70/kantin-backend/pkg/storage/persistence"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var data model.UserRegister
	c.BindJSON(&data)

	db, status := c.MustGet("db").(storage.Database)
	if !status {
    fmt.Println(status)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": status,
		})
		return
	}

	res, err := persistence.UserRegister(&data, db)
	if err != nil {
    fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})

		return
	}

	// TODO: Should return token, still missing middleware
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"token":   res,
	})
}


func UserLogin(c *gin.Context) {
	db, status := c.MustGet("db").(storage.Database)
	if status {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": status,
		})
		return
	}

	var data model.UserLogin
	c.BindJSON(&data)

	res, err := persistence.UserLogin(&data, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	// TODO: Should return token, still missing middleware
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"token":   res,
	})
}

func UserGetAccount(c *gin.Context) {
	db, status := c.MustGet("db").(storage.Database)
	if !status {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": status,
		})
		return
	}

	user_id := c.Param("user_id")
	res, err := persistence.UserGetAccount(user_id, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        res.ID,
		"email":     res.Email,
		"firstname": res.Firstname,
		"lastname":  res.Lastname,
	})
}

func UserUpdateAccount(c *gin.Context) {
	db, status := c.MustGet("db").(storage.Database)
	if !status {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": status,
		})
		return
	}

	var data model.User
	c.BindJSON(&data)

	res, err := persistence.UserUpdateAccount(&data, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        res.ID,
		"email":     res.Email,
		"firstname": res.Firstname,
		"lastname":  res.Lastname,
	})
}

func UserGetCategory(c *gin.Context)   {} // Marked
func UserGetSeller(c *gin.Context)     {} // Marked
func UserGetSellerFood(c *gin.Context) {} // Marked
func UserGetFoodDetail(c *gin.Context) {} // Marked

func UserMakeOrder(c *gin.Context) {
 //  var data model.Order
	// db, status := c.MustGet("db").(storage.Database)
	// if !status {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": status,
	// 	})
	// 	return
	// }
}

func UserGetOrderStatus(c *gin.Context) {
	// db, status := c.MustGet("db").(storage.Database)
	// if !status {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": status,
	// 	})
	// 	return
	// }
}

func UserRequestOrderCancel(c *gin.Context) {
	// db, status := c.MustGet("db").(storage.Database)
	// if !status {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": status,
	// 	})
	// 	return
	// }
}

func UserGetOrders(c *gin.Context) {
	// db, status := c.MustGet("db").(storage.Database)
	// if !status {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": status,
	// 	})
	// 	return
	// }
}
