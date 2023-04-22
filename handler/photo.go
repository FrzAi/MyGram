package handler

import (
	"Final_Project/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h HttpServer) PhotoGetAll(c *gin.Context) {

	res, err := h.app.PhotoGetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h HttpServer) PhotoGet(c *gin.Context) {

	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	var photo model.Photo

	photo.ID = id

	res, err := h.app.PhotoGet(photo)

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

func (h HttpServer) PhotoCreate(c *gin.Context) {
	var newPhoto model.Photo

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))

	if err := c.ShouldBindJSON(&newPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	newPhoto.UserID = int(userID)

	res, err := h.app.PhotoCreate(newPhoto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h HttpServer) PhotoUpdate(c *gin.Context) {
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
	var photo model.Photo

	photo.ID = id

	photo, err = h.app.PhotoAuthorization(photo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})

		return
	}

	if photo.UserID != userID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "you're not allowed to access this",
		})

		return
	}

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	photo.ID = id

	res, err := h.app.PhotoUpdate(photo)

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

func (h HttpServer) PhotoDelete(c *gin.Context) {
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
	var photo model.Photo

	photo.ID = id

	photo, err = h.app.PhotoAuthorization(photo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})

		return
	}

	if photo.UserID != userID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "you're not allowed to access this",
		})

		return
	}

	photo.ID = id

	err = h.app.PhotoDelete(photo)

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
		"message": "Photo deleted successfully",
	})
}
