package controllers

import (
	"net/http"
	"web-api/internal/api/services"
	"web-api/internal/pkg/models/request"
	"web-api/internal/pkg/models/response"

	"github.com/gin-gonic/gin"
)

// CONTROLLER - Login
type LoginController struct {
	*BaseController
}

var LG = &LoginController{}

// func - Login
func (c *LoginController) Login(ctx *gin.Context) {
	var requestParams request.LoginRequest
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	// Gọi dịch vụ đăng nhập
	service := services.LoginService{}
	result, err := service.Login(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	// Trả về thông tin người dùng và token
	response.OkWithData(ctx, result)
}
