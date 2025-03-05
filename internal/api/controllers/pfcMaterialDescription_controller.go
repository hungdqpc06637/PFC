package controllers

import (
	"log"
	"net/http"
	"web-api/internal/api/services"
	"web-api/internal/pkg/models/response"
	"web-api/internal/pkg/models/types"

	"github.com/gin-gonic/gin"
)

type PFCMaterialDescriptionController struct {
	*BaseController
}

var PFCMaterialDescription = &PFCMaterialDescriptionController{}

func (c *PFCMaterialDescriptionController) InsertPFCMaterialDescription(ctx *gin.Context) {
	var requestParams types.PFC_MaterialDescription
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCMaterialDescription.InsertNewPFCMaterialDescription(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCMaterialDescriptionController) UpdatePFCMaterialDescription(ctx *gin.Context) {
	var requestParams types.PFC_MaterialDescription
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCMaterialDescription.UpdatePFCMaterialDescription(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCMaterialDescriptionController) DeletePFCMaterialDescription(ctx *gin.Context) {
	var requestParams types.PFC_MaterialDescription
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCMaterialDescription.DeletePFCMaterialDescription(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCMaterialDescriptionController) DeletePFCMaterialDescriptionByModelID(ctx *gin.Context) {
	var requestParams types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCMaterialDescription.DeletePFCMaterialDescriptionByModelID(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
