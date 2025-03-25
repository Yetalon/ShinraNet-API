package controllers

import (
	"Backend/database"
	"Backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid post data")
		return
	}
	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to create post")
	}
	c.JSON(http.StatusOK, post.Title)
}

func GetPostsByTitle(c *gin.Context) {
	title := c.DefaultQuery("title", "")

	var Posts []models.Post
	if title == "" {
		if err := database.DB.Find(&Posts).Error; err != nil {
			c.JSON(http.StatusInternalServerError, "failed to retrive posts")
			return
		}
		c.JSON(http.StatusOK, Posts)
		return
	} else {
		if err := database.DB.Where("title = ?", title).First(&Posts).Error; err != nil {
			c.JSON(http.StatusOK, Posts)
			return
		}
	}
	c.JSON(http.StatusOK, Posts)
}

func GetPostByUserId(c *gin.Context) {
	id := c.Param("id")
	userid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	var Posts []models.Post
	if err := database.DB.Where("user_id = ?", userid).Find(&Posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to find posts")
		return
	}
	c.JSON(http.StatusOK, Posts)
}

func GetPostById(c *gin.Context) {
	id := c.Param("id")
	postid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	var Post models.Post
	if err := database.DB.Where("id = ?", postid).First(&Post).Error; err != nil {
		c.JSON(http.StatusBadRequest, "post not found")
		return
	}
	c.JSON(http.StatusOK, Post)
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	postid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid post id")
		return
	}
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, "invalid user data")
		return
	}
	var existingpost models.Post
	if err := database.DB.Where("id = ?", postid).First(&existingpost).Error; err != nil {
		c.JSON(http.StatusBadRequest, "post does not exist")
		return
	}
	existingpost.Title = post.Title
	existingpost.PostText = post.PostText
	if err := database.DB.Save(&existingpost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, "could not save post update")
		return
	}
	c.JSON(http.StatusOK, "post has been updated successfully")
}

func PatchPostTitle(c *gin.Context) {
	id := c.Param("id")
	postid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid post id")
		return
	}
	title := c.DefaultQuery("name", "")
	if title == "" {
		c.String(http.StatusBadRequest, "invalid id parameter")
		return
	}
	if err := database.DB.Model(&models.Post{}).Where("id = ?", postid).Update("title", title).Error; err != nil {
		c.JSON(http.StatusInternalServerError, "could not update")
		return
	}
	c.JSON(http.StatusOK, "post has been patched successfully")
}

func PatchPostText(c *gin.Context) {
	id := c.Param("id")
	postid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid post id")
		return
	}
	postText := c.DefaultQuery("post_text", "")
	if postText == "" {
		c.String(http.StatusBadRequest, "post text is empty")
		return
	}
	if err := database.DB.Model(&models.Post{}).Where("id = ?", postid).Update("post_text", postText).Error; err != nil {
		c.String(http.StatusInternalServerError, "could not update")
		return
	}
	c.JSON(http.StatusOK, "post has been patched successfully")
}

func IncrementPostLikes(c *gin.Context) {
	id := c.Param("id")
	postid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid post id")
		return
	}
	var post models.Post
	if err := database.DB.Where("id = ?", postid).First(&post); err != nil {
		c.String(http.StatusBadRequest, "post does not exist")
		return
	}
	if err := database.DB.Model(&models.Post{}).Where("id = ?", postid).Update("likes", post.Likes+1).Error; err != nil {
		c.String(http.StatusInternalServerError, "could not update likes")
		return
	}
	c.String(http.StatusOK, "incremented post likes successfully")
}

func DecrementPostLikes(c *gin.Context) {
	id := c.Param("id")
	postid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid post id")
		return
	}
	var post models.Post
	if err := database.DB.Where("id = ?", postid).First(&post); err != nil {
		c.String(http.StatusBadRequest, "post does not exist")
		return
	}
	if err := database.DB.Model(&models.Post{}).Where("id = ?", postid).Update("likes", post.Likes-1).Error; err != nil {
		c.String(http.StatusInternalServerError, "could not update likes")
		return
	}
	c.String(http.StatusOK, "decremented post likes successfully")
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	postid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid post id")
		return
	}
	result := database.DB.Delete(&models.Post{}, postid)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, "could not delete post")
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, "post not found")
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("post %d has been deleted", postid))
}
