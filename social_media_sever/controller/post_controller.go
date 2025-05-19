package controllers

import (
	"net/http"
	"social_media_sever/config"
	"social_media_sever/models"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

func NewPostController() *PostController {
	return &PostController{}
}

func (ctl *PostController) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&post)
	c.JSON(http.StatusOK, post)
}

func (ctl *PostController) GetPosts(c *gin.Context) {
	var posts []models.Post
	config.DB.Preload("Comments").Find(&posts)
	c.JSON(http.StatusOK, posts)
}

func (ctl *PostController) GetPostById(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	result := config.DB.Preload("Comments").First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (ctl *PostController) UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	// Check if post exists
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Bind JSON to post
	var updatedPost models.Post
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	post.Title = updatedPost.Title
	post.Content = updatedPost.Content

	config.DB.Save(&post)
	c.JSON(http.StatusOK, post)
}

func (ctl *PostController) DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	// Check if post exists
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	config.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
