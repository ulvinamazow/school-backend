package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ulvinamazow/studentAPI/config"
	"github.com/ulvinamazow/studentAPI/models"
)

func ListUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(http.StatusOK, &users)

}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, &user)
	}
}
