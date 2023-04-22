package handler

import (
	"Final_Project/helper"
	"Final_Project/model"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) UserRegister(c *gin.Context) {
	contentType := helper.GetContentType(c)
	user := model.User{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	if user.Age < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Maaf, Anda Belum cukup umur!",
			"status":  http.StatusBadRequest,
		})
		return
	}

	if len(user.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password harus lebih dari 6 karakter!",
			"status":  http.StatusBadRequest,
		})
		return
	}

	_, err := mail.ParseAddress(user.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email tidak valid!",
			"status":  http.StatusBadRequest,
		})
		return
	}

	res, err := h.app.UserRegister(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h HttpServer) UserLogin(c *gin.Context) {
	contentType := helper.GetContentType(c)
	user := model.User{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	if len(user.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password harus lebih dari 6 karakter!",
			"status":  http.StatusBadRequest,
		})
	}

	passwordClient := user.Password

	res, err := h.app.UserLogin(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "invalid email or password",
		})
		return
	}

	isValid := helper.ComparePass([]byte(res.Password), []byte(passwordClient))
	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "invalid email or password",
		})
		return
	}

	token := helper.GenerateToken(uint(res.ID), res.Username)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
