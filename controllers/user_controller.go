package controllers

import (
	"net/http"
	"strconv"

	"profile-picture-api/app"
	"profile-picture-api/database"
	"profile-picture-api/helpers"
	"profile-picture-api/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User

	database.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"status":  "Succeed",
		"message": "Fetch all users",
		"data":    users,
	})
}

func GetUserbyID(c *gin.Context) {
	var user models.User

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "ID is not valid",
		})
		return
	}

	if err := database.DB.First(&user, "id =?", userID).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "Failed",
				"message": "Data not found",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "Failed",
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Succeed",
		"message": "Fetch a user",
		"data":    user,
	})

}

func UpdateUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "ID is not valid",
		})
		return
	}

	var userReq app.UserUpdate

	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status":  "Failed",
			"message": err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(userReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": err.Error(),
		})
		return
	}

	hashPassword, err := helpers.HashPassword(userReq.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": err.Error(),
		})
		return
	}
	userReq.Password = hashPassword

	userModel := models.User{
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: userReq.Password,
	}

	if database.DB.Model(&userModel).Where("id = ?", userID).Updates(&userModel).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update data",
			"status":  "Failed",
		})
		return
	}

	c.Set("user", nil)

	c.JSON(http.StatusOK, gin.H{
		"status":  "Succeed",
		"message": "Success update data",
	})

}

func DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "ID is not valid",
		})
		return
	}

	var user models.User

	if database.DB.Unscoped().Delete(&user, userID).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Failed to delete data",
		})
		return
	}

	c.Set("user", nil)
	c.SetCookie("Authorization", "", -1, "", "", true, true)

	c.JSON(200, gin.H{
		"status":  "Succeed",
		"message": "Success delete data",
	})
}
