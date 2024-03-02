package controllers

import (
	"net/http"
	"os"
	"time"

	"profile-picture-api/app"
	"profile-picture-api/database"
	"profile-picture-api/helpers"
	"profile-picture-api/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Register(c *gin.Context) {
	var user app.UserRegister

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": err.Error(),
		})
		return
	}
	hashPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": err.Error(),
		})
		return
	}

	user.Password = hashPassword

	userModel := models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := database.DB.Create(&userModel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Succeed",
		"message": "User created",
	})

}

func Login(c *gin.Context) {
	var userReq app.UserLogin

	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
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

	var data models.User

	if err := database.DB.First(&data, "email = ?", userReq.Email).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Failed",
			"message": "Invalid email or password",
		})
		return
	}

	if err := helpers.CheckHashedPassword(userReq.Password, data.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "Failed",
			"message": "Invalid email or password",
		})
		return
	}

	expTime := time.Now().Add(10 * time.Minute)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": expTime.Unix(),
		"sub": data.ID,
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "Failed",
			"message": err.Error(),
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenStr, 3600*24, "", "", true, true)

	c.JSON(http.StatusAccepted, gin.H{
		"status":  "Succeed",
		"message": "Login Success",
	})

}

func Logout(c *gin.Context) {
	c.Set("user", nil)
	c.SetCookie("Authorization", "", -1, "", "", true, true)

	c.JSON(http.StatusOK, gin.H{
		"status":  "Succeed",
		"message": "Logout Success",
	})
}
