package controllers

import (
	"net/http"
	"social_media_sever/config"
	"social_media_sever/models"

	"github.com/gin-gonic/gin"
)

type CommentController struct{}

func NewCommentController() *CommentController {
	return &CommentController{}
}

func (ctl *CommentController) CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&comment)
	c.JSON(http.StatusOK, comment)
}

func (ctl *CommentController) GetComments(c *gin.Context) {
	var comments []models.Comment
	config.DB.Find(&comments)
	c.JSON(http.StatusOK, comments)
}

func (ctl *CommentController) GetCommentsByPostId(c *gin.Context) {
	postID := c.Param("id")
	var comments []models.Comment

	result := config.DB.Where("post_id = ?", postID).Find(&comments)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (ctl *CommentController) UpdateComment(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment

	// Check if comment exists
	if err := config.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Bind JSON to comment
	var updatedComment models.Comment
	if err := c.ShouldBindJSON(&updatedComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update field
	comment.Content = updatedComment.Content

	config.DB.Save(&comment)
	c.JSON(http.StatusOK, comment)
}

func (ctl *CommentController) DeleteComment(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment

	// Check if comment exists
	if err := config.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	config.DB.Delete(&comment)
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
