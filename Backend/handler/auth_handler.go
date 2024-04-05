package handler

import (
	"backend/entity"
	"backend/errorHandler"
	"backend/helper"
	"backend/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var register entity.User

	if err := c.ShouldBindJSON(&register); err != nil {
		errorHandler.HandleError(c, &errorHandler.Error400{Message: err.Error()})
	}

	result, err := h.service.Register(&register)
	if err != nil {

		errorHandler.HandleError(c, err)
		return
	}
	res := &helper.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "User Registered Succesfully",
		Data:       result,
	}

	c.JSON(http.StatusCreated, res)
}

func (h *authHandler) Login(c *gin.Context) {
	var login entity.LoginRequest
	err := c.ShouldBindJSON(&login)

	if err != nil {
		errorHandler.HandleError(c, &errorHandler.Error400{Message: err.Error()})
		return
	}
	result, err := h.service.Login(&login)
	if err != nil {
		errorHandler.HandleError(c, err)
		return
	}

	res := helper.Response(helper.ResponseParams{
		Message: "Login Succesfully",
		Data:    result,
	})

	c.Header("Authorization", result.AccessToken)

	c.JSON(http.StatusOK, res)
	auth := c.Writer.Header().Get("Authorization")
	log.Print(auth)

}
