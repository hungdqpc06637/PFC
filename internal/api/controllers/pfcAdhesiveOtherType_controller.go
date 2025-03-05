package controllers

import (
	"log"
	"net/http"
	"web-api/internal/api/services"
	"web-api/internal/pkg/models/response"
	"web-api/internal/pkg/models/types"

	"github.com/gin-gonic/gin"
)

type PFCAdhesiveOtherTypeController struct {
	*BaseController
}

var PFCAdhesiveOtherType = &PFCAdhesiveOtherTypeController{}

func (c *PFCAdhesiveOtherTypeController) InsertPFCAdhesiveOtherType(ctx *gin.Context) {
	var requestParams types.PFC_AdhesiveOtherType
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCAdhesiveOtherType.InsertNewPFCAdhesiveOtherType(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCAdhesiveOtherTypeController) UpdatePFCAdhesiveOtherType(ctx *gin.Context) {
	var requestParams types.PFC_AdhesiveOtherType
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCAdhesiveOtherType.UpdatePFCAdhesiveOtherType(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCAdhesiveOtherTypeController) DeletePFCAdhesiveOtherType(ctx *gin.Context) {
	var requestParams types.PFC_AdhesiveOtherType
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCAdhesiveOtherType.DeletePFCAdhesiveOtherType(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCAdhesiveOtherTypeController) DeletePFCAdhesiveOtherTypeByModelID(ctx *gin.Context) {
	var requestParams types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCAdhesiveOtherType.DeletePFCAdhesiveOtherTypeByModelID(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
