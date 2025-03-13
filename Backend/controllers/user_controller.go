package controllers

import (
	"Backend/database"
	"Backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid user data")
		return
	}
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to create user")
		return
	}
	c.JSON(http.StatusOK, user.Username)
}

func ParseUserId(id string) (int, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

// Get request by userID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	userID, err := ParseUserId(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "User id not found")
		return
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// get request by userName
func GetUserByName(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	var ListOfUsers []models.User
	if name == "" {
		if err := database.DB.Find(&ListOfUsers).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch users"})
			return
		}
		c.JSON(http.StatusOK, ListOfUsers)
		return
	}
	var user models.User
	if err := database.DB.Where("username = ?", name).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, []models.User{})
		return
	}
	ListOfUsers = append(ListOfUsers, user)
	c.JSON(http.StatusOK, ListOfUsers)
}

// update for user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	userID, err := ParseUserId(id)
	if err != nil {
		c.String(http.StatusNotFound, "User id not found")
		return
	}

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.String(http.StatusBadRequest, "no user")
		return
	}

	var existingUser models.User
	if err := database.DB.First(&existingUser, userID).Error; err != nil {
		c.String(http.StatusBadRequest, "user does not exist")
	}

	existingUser.Username = updatedUser.Username
	existingUser.ImageURL = updatedUser.ImageURL

	if err := database.DB.Save(&existingUser).Error; err != nil {
		c.String(http.StatusInternalServerError, "could not update")
		return
	}
	c.String(http.StatusOK, "user has been updated")
}

// Patch userName
func PatchUserName(c *gin.Context) {
	id := c.Param("id")

	userID, err := ParseUserId(id)
	if err != nil {
		c.String(http.StatusNotFound, "User id not found")
		return
	}
	name := c.DefaultQuery("name", "")
	if name == "" {
		c.String(http.StatusBadRequest, "invalid parameter")
	}

	if err := database.DB.Model(&models.User{}).Where("id = ?", userID).Update("username", name).Error; err != nil {
		c.String(http.StatusInternalServerError, "could not update")
	}
	c.JSON(http.StatusOK, "username has been updated")
}

// Patch userAdmin
func PatchAdmin(c *gin.Context) {
	id := c.Param("id")

	userID, err := ParseUserId(id)
	if err != nil {
		c.String(http.StatusNotFound, "User id not found")
		return
	}

	isAdmin := c.DefaultQuery("is_admin", "")
	var userAdminStatus bool
	if userAdminStatus, err = strconv.ParseBool(isAdmin); err != nil {
		c.String(http.StatusBadRequest, "invalid admin status")
	}

	if err := database.DB.Model(&models.User{}).Where("id = ?", userID).Update("is_admin", userAdminStatus).Error; err != nil {
		c.String(http.StatusInternalServerError, "could not update")
		return
	}
	c.JSON(http.StatusOK, "admin status has been updated")
}

// Patch userImage
func PatchImage(c *gin.Context) {
	id := c.Param("id")

	userID, err := ParseUserId(id)
	if err != nil {
		c.String(http.StatusNotFound, "User id not found")
		return
	}

	imageURL := c.DefaultQuery("image_url", "")
	if imageURL == "" {
		c.String(http.StatusNotFound, "No image url found")
		return
	}

	if err := database.DB.Model(&models.User{}).Where("id = ?", userID).Update("image_url", imageURL).Error; err != nil {
		c.String(http.StatusInternalServerError, "could not update")
		return
	}
	c.JSON(http.StatusOK, "Image url has been update")
}

// Delete User
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	userID, err := ParseUserId(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "User id not found")
		return
	}

	result := database.DB.Delete(&models.User{}, userID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, "could not delete user")
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, "user not found")
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("user %s has been deleted", userID))
}
