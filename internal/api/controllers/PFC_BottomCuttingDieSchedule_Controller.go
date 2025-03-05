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

type PFCBottomCuttingDieSchedule struct {
	*BaseController
}

var PFCBottomCuttingDieSchedul = &PFCBottomCuttingDieSchedule{}

func (c *PFCBottomCuttingDieSchedule) GetAllPFCBottomCuttingDieSchedule(ctx *gin.Context) {
	var requestParams *types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCBottomCuttingDieSchedul.GetAllPFCBottomCuttingDieSchedule(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCBottomCuttingDieSchedule) InsertPFCBottomCuttingDieSchedule(ctx *gin.Context) {
	var requestParams types.PFC_BottomCuttingDieSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCBottomCuttingDieSchedul.InsertNewPFCBottomCuttingDieSchedule(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCBottomCuttingDieSchedule) UpdatePFCBottomCuttingDieSchedule(ctx *gin.Context) {
	var requestParams types.PFC_BottomCuttingDieSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCBottomCuttingDieSchedul.UpdatePFCBottomCuttingDieSchedule(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCBottomCuttingDieSchedule) DeletePFCBottomCuttingDieSchedule(ctx *gin.Context) {
	var requestParams types.PFC_BottomCuttingDieSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCBottomCuttingDieSchedul.DeletePFCBottomCuttingDieSchedule(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

// //ITEM PFC Skiving Instructions
func (c *PFCBottomCuttingDieSchedule) InsertPFCItemBottomCuttingDieSchedule(ctx *gin.Context) {
	var requestParams types.PFC_ItemBottomCuttingDieSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}
	result, err := services.PFCBottomCuttingDieSchedul.InsertNewPFCItemBottomCuttingDieSchedule(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCBottomCuttingDieSchedule) GetAllPFCItemBottomCuttingDieSchedule(ctx *gin.Context) {
	var requestParams *types.PFC_BottomCuttingDieSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCBottomCuttingDieSchedul.GetAllPFCItemBottomCuttingDieSchedule(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCBottomCuttingDieSchedule) UpdatePFCItemBottomCuttingDieSchedule(ctx *gin.Context) {
	var requestParams types.PFC_ItemBottomCuttingDieSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCBottomCuttingDieSchedul.UpdatePFCItemBottomCuttingDieSchedule(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCBottomCuttingDieSchedule) DeletePFCItemBottomCuttingDieSchedule(ctx *gin.Context) {
	var requestParams types.PFC_ItemBottomCuttingDieSchedule
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCBottomCuttingDieSchedul.DeletePFCItemBottomCuttingDieSchedule(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
