package controllers

import (
	"log"
	"net/http"
	"web-api/internal/api/services"
	"web-api/internal/pkg/models/response"

	"web-api/internal/pkg/models/types"

	"github.com/gin-gonic/gin"
)

type PFCModelController struct {
	*BaseController
}

var PFCModel = &PFCModelController{}

func (c *PFCModelController) GetAllPFCModel(ctx *gin.Context) {
	result, err := services.PFCModel.GetAllPFCModel()

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCModelController) InsertNewPFCModel(ctx *gin.Context) {
	var requestParams types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCModel.InsertNewModel(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCModelController) UpdatePFCModel(ctx *gin.Context) {
	var requestParams types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCModel.UpdatePFCModel(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCModelController) DeletePFCModel(ctx *gin.Context) {
	var requestParams types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCModel.DeletePFCModel(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
