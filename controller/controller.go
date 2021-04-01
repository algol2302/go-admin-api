package controller

import (
	"github.com/algol2302/go-admin-api/config"
	"github.com/algol2302/go-admin-api/model"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterEndPoint(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userCheck model.User
	config.GetDB().First(&userCheck, "username = ?", user.Username)

	if userCheck.ID > 0 {
		c.JSON(http.StatusConflict, gin.H{"message": "User already exists"})
		return
	}

	config.GetDB().Save(&user)

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func FetchAllUsers(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	var user model.User
	config.GetDB().Where("id = ?", claims[config.IdentityKey]).First(&user)

	if user.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	var users []model.User
	config.GetDB().Order("created_at desc").Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found!", "data": users})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func FetchSingleUser(c *gin.Context) {
	userID := c.Param("id")

	if len(userID) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	var user model.User
	config.GetDB().First(&user, userID)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found!"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	if len(userID) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	var newUser model.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	config.GetDB().First(&user, userID)

	if user.ID <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No user found!"})
		return
	}

	config.GetDB().Model(&user).Update("username", newUser.Username)
	config.GetDB().Model(&user).Update("email", newUser.Email)
	config.GetDB().Model(&user).Update("role", newUser.Role)
	config.GetDB().Model(&user).Update("password", newUser.Password)

	config.GetDB().First(&user, userID)

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully!", "user": user})
}

func DeleteUser(c *gin.Context) {
	var user model.User
	userID := c.Param("id")

	if len(userID) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	config.GetDB().First(&user, userID)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No user found!"})
		return
	}

	config.GetDB().Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully!", "user": user})
}
