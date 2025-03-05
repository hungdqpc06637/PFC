package controllers

import (
	"log"
	"net/http"
	"web-api/internal/api/services"
	"web-api/internal/pkg/models/response"
	"web-api/internal/pkg/models/types"

	"github.com/gin-gonic/gin"
)

type PFCItemUpperCuttingDieScheduleController struct {
	*BaseController
}

var PFCItemUpperCuttingDieSchedule = &PFCItemUpperCuttingDieScheduleController{}

func (c *PFCItemUpperCuttingDieScheduleController) GetAllPFCItemUpperCuttingDieScheduleByModelID(ctx *gin.Context) {
	var requestParams types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCItemUpperCuttingDieSchedule.GetAllItemUpperCuttingDieScheduleByModelID(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCItemUpperCuttingDieScheduleController) GetAllPFCItemUpperCuttingDieSchedule(ctx *gin.Context) {
	var requestParams types.PFC_UpperCuttingSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCItemUpperCuttingDieSchedule.GetAllItemUpperCuttingDieSchedule(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCItemUpperCuttingDieScheduleController) InsertPFCItemUpperCuttingDieSchedule(ctx *gin.Context) {
	var requestParams types.PFC_ItemUpperCuttingSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCItemUpperCuttingDieSchedule.InsertNewPFCItemUpperCuttingDieSchedule(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCItemUpperCuttingDieScheduleController) UpdatePFCItemUpperCuttingDieSchedule(ctx *gin.Context) {
	var requestParams types.PFC_ItemUpperCuttingSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCItemUpperCuttingDieSchedule.UpdatePFCItemUpperCuttingDieSchedule(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCItemUpperCuttingDieScheduleController) DeletePFCItemUpperCuttingDieSchedule(ctx *gin.Context) {
	var requestParams types.PFC_ItemUpperCuttingSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCItemUpperCuttingDieSchedule.DeletePFCItemUpperCuttingDieSchedule(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCItemUpperCuttingDieScheduleController) DeletePFCItemUpperCuttingDieScheduleByModelID(ctx *gin.Context) {
	var requestParams types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCItemUpperCuttingDieSchedule.DeletePFCItemUpperCuttingDieScheduleByModelID(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
