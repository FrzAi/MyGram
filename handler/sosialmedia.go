package handler

import (
	"Final_Project/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h HttpServer) SocialMediaGetAll(c *gin.Context) {

	res, err := h.app.SocialMediaGetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h HttpServer) SocialMediaGet(c *gin.Context) {

	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	var photo model.SocialMedia

	photo.ID = id

	res, err := h.app.SocialMediaGet(photo)

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

func (h HttpServer) SocialMediaCreate(c *gin.Context) {
	var newSocialMedia model.SocialMedia

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))

	if err := c.ShouldBindJSON(&newSocialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	newSocialMedia.UserID = int(userID)

	res, err := h.app.SocialMediaCreate(newSocialMedia)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"status":  http.StatusBadRequest,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"status":  http.StatusInternalServerError,
			})
		}

		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h HttpServer) SocialMediaUpdate(c *gin.Context) {
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
	var photo model.SocialMedia

	photo.ID = id

	photo, err = h.app.SocialMediaAuthorization(photo)

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

	res, err := h.app.SocialMediaUpdate(photo)

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

func (h HttpServer) SocialMediaDelete(c *gin.Context) {
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
	var photo model.SocialMedia

	photo.ID = id

	photo, err = h.app.SocialMediaAuthorization(photo)

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

	err = h.app.SocialMediaDelete(photo)

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
		"message": "Social Media deleted successfully",
	})
}
