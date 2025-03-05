package controllers

import (
	"log"
	"net/http"
	"web-api/internal/api/services"
	"web-api/internal/pkg/models/response"
	"web-api/internal/pkg/models/types"

	"github.com/gin-gonic/gin"
)

type PFCUpperCuttingDieScheduleController struct {
	*BaseController
}

var PFCUpperCuttingDieSchedule = &PFCUpperCuttingDieScheduleController{}

func (c *PFCUpperCuttingDieScheduleController) GetAllPFCUpperCuttingDieSchedule(ctx *gin.Context) {
	var requestParams types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCUpperCuttingDieSchedule.GetAllPFCUpperCuttingDieSchedule(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCUpperCuttingDieScheduleController) InsertPFCUpperCuttingDieSchedule(ctx *gin.Context) {
	var requestParams types.PFC_UpperCuttingSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCUpperCuttingDieSchedule.InsertNewPFCUpperCuttingDieSchedule(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCUpperCuttingDieScheduleController) UpdatePFCUpperCuttingDieSchedule(ctx *gin.Context) {
	var requestParams types.PFC_UpperCuttingSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCUpperCuttingDieSchedule.UpdatePFCUpperCuttingDieSchedule(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCUpperCuttingDieScheduleController) DeletePFCUpperCuttingDieSchedule(ctx *gin.Context) {
	var requestParams types.PFC_UpperCuttingSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCUpperCuttingDieSchedule.DeletePFCUpperCuttingDieSchedule(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCUpperCuttingDieScheduleController) DeletePFCUpperCuttingDieScheduleByModelID(ctx *gin.Context) {
	var requestParams types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCUpperCuttingDieSchedule.DeletePFCUpperCuttingDieScheduleByModelID(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
