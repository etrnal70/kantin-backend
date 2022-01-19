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

type UserHandler struct {
	Conn storage.Database
}

func (db *UserHandler) UserRegister(c *gin.Context) {
	var data model.UserRegister
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	res, err := persistence.UserRegister(&data, db.Conn)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Unable to register user :%s", err),
		})
	}

	token, err := middleware.GenerateToken(res.ID.String())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    res,
	})
}

func (db *UserHandler) UserLogin(c *gin.Context) {
	var data model.UserLogin
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	id, err := persistence.UserLogin(&data, db.Conn)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
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

func (db *UserHandler) UserGetAccount(c *gin.Context) {
	user_id, _ := uuid.FromString(c.Param("user_id"))
	res, err := persistence.UserGetAccount(user_id, db.Conn)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        res.ID,
		"email":     res.Email,
		"firstname": res.Firstname,
		"lastname":  res.Lastname,
	})
}

func (db *UserHandler) UserUpdateAccount(c *gin.Context) {
	user_id, _ := uuid.FromString(c.Param("user_id"))
	var data model.User
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	// TODO: Check if password change

	res, err := persistence.UserUpdateAccount(user_id, &data, db.Conn)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        res.ID,
		"email":     res.Email,
		"firstname": res.Firstname,
		"lastname":  res.Lastname,
	})
}

func (db *UserHandler) UserMakeOrder(c *gin.Context) {
	//  var data model.Order
	// db, status := c.MustGet("db").(storage.Database)
	// if !status {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": status,
	// 	})
	// 	return
	// }
}

func (db *UserHandler) UserGetOrderStatus(c *gin.Context) {
	// db, status := c.MustGet("db").(storage.Database)
	// if !status {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": status,
	// 	})
	// 	return
	// }
}

func (db *UserHandler) UserRequestOrderCancel(c *gin.Context) {
	// db, status := c.MustGet("db").(storage.Database)
	// if !status {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": status,
	// 	})
	// 	return
	// }
}

func (db *UserHandler) UserGetOrders(c *gin.Context) {
	// db, status := c.MustGet("db").(storage.Database)
	// if !status {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": status,
	// 	})
	// 	return
	// }
}
