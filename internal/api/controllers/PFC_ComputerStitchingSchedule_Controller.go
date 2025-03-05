//

package controllers

import (
	"log"
	"net/http"
	"web-api/internal/api/services"
	"web-api/internal/pkg/models/response"
	"web-api/internal/pkg/models/types"

	"github.com/gin-gonic/gin"
)

type PFCComputerStitchingSchedule struct {
	*BaseController
}

var PFCComputerStitchingSchedu = &PFCComputerStitchingSchedule{}

func (c *PFCComputerStitchingSchedule) GetAllPFCComputerStitchingSchedule(ctx *gin.Context) {
	var requestParams *types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCComputerStitchingSchedu.GetAllPFCComputerStitchingSchedule(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCComputerStitchingSchedule) InsertPFCComputerStitchingSchedule(ctx *gin.Context) {
	var requestParams types.PFC_ComputerStitchingSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCComputerStitchingSchedu.InsertNewPFCComputerStitchingSchedule(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCComputerStitchingSchedule) UpdatePFCComputerStitchingSchedule(ctx *gin.Context) {
	var requestParams types.PFC_ComputerStitchingSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCComputerStitchingSchedu.UpdatePFCComputerStitchingSchedule(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCComputerStitchingSchedule) DeletePFCComputerStitchingSchedule(ctx *gin.Context) {
	var requestParams types.PFC_ComputerStitchingSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCComputerStitchingSchedu.DeletePFCComputerStitchingSchedule(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

// //ITEM PFC Skiving Instructions
func (c *PFCComputerStitchingSchedule) InsertPFCItemComputerStitchingSchedule(ctx *gin.Context) {
	var requestParams types.PFC_ItemComputerStitchingSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}
	result, err := services.PFCComputerStitchingSchedu.InsertNewPFCItemComputerStitchingSchedule(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCComputerStitchingSchedule) GetAllPFCItemComputerStitchingSchedule(ctx *gin.Context) {
	var requestParams *types.PFC_ComputerStitchingSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCComputerStitchingSchedu.GetAllPFCItemComputerStitchingSchedule(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCComputerStitchingSchedule) UpdatePFCItemComputerStitchingSchedule(ctx *gin.Context) {
	var requestParams types.PFC_ItemComputerStitchingSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCComputerStitchingSchedu.UpdatePFCItemComputerStitchingSchedule(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCComputerStitchingSchedule) DeletePFCItemComputerStitchingSchedule(ctx *gin.Context) {
	var requestParams types.PFC_ItemComputerStitchingSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCComputerStitchingSchedu.DeletePFCItemComputerStitchingSchedule(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
