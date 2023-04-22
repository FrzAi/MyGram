package handler

import (
	"Final_Project/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h HttpServer) CommentGetAll(c *gin.Context) {

	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	var comment model.Comment

	comment.ID = id

	res, err := h.app.CommentGetAll(comment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h HttpServer) CommentGet(c *gin.Context) {

	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	var comment model.Comment

	comment.ID = id

	res, err := h.app.CommentGet(comment)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
				"status":  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h HttpServer) CommentCreate(c *gin.Context) {

	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	var newComment model.Comment

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))

	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	newComment.UserID = int(userID)
	newComment.PhotoID = id

	res, err := h.app.CommentCreate(newComment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h HttpServer) CommentUpdate(c *gin.Context) {
	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))
	var comment model.Comment

	comment.ID = id

	comment, err = h.app.CommentAuthorization(comment)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})

		return
	}

	if comment.UserID != userID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "you're not allowed to access this",
		})

		return
	}

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	comment.ID = id

	res, err := h.app.CommentUpdate(comment)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
				"status":  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h HttpServer) CommentDelete(c *gin.Context) {
	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))
	var comment model.Comment

	comment.ID = id

	comment, err = h.app.CommentAuthorization(comment)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})

		return
	}

	if comment.UserID != userID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "you're not allowed to access this",
		})

		return
	}

	comment.ID = id

	err = h.app.CommentDelete(comment)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
				"status":  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Comment deleted successfully",
	})
}
