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

type PFCBottomSilkScreenProcess struct {
	*BaseController
}

var PFCBottomSilkScreenProces = &PFCBottomSilkScreenProcess{}

func (c *PFCBottomSilkScreenProcess) GetAllPFCBottomSilkScreenProcess(ctx *gin.Context) {
	var requestParams *types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCBottomSilkScreenProces.GetAllPFCBottomSilkScreenProcess(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCBottomSilkScreenProcess) InsertPFCBottomSilkScreenProcess(ctx *gin.Context) {
	var requestParams types.PFC_BottomSilkScreenProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCBottomSilkScreenProces.InsertNewPFCBottomSilkScreenProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCBottomSilkScreenProcess) UpdatePFCBottomSilkScreenProcess(ctx *gin.Context) {
	var requestParams types.PFC_BottomSilkScreenProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCBottomSilkScreenProces.UpdatePFCBottomSilkScreenProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCBottomSilkScreenProcess) DeletePFCBottomSilkScreenProcess(ctx *gin.Context) {
	var requestParams types.PFC_BottomSilkScreenProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCBottomSilkScreenProces.DeletePFCBottomSilkScreenProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

// //ITEM PFC Bottom SilkScreen Process
func (c *PFCBottomSilkScreenProcess) InsertPFCItemBottomSilkScreenProcess(ctx *gin.Context) {
	var requestParams types.PFC_ItemBottomSilkScreenProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}
	result, err := services.PFCBottomSilkScreenProces.InsertNewPFCItemBottomSilkScreenProcess(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCBottomSilkScreenProcess) GetAllPFCItemBottomSilkScreenProcess(ctx *gin.Context) {
	var requestParams *types.PFC_BottomSilkScreenProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCBottomSilkScreenProces.GetAllPFCItemBottomSilkScreenProcess(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCBottomSilkScreenProcess) UpdatePFCItemBottomSilkScreenProcess(ctx *gin.Context) {
	var requestParams types.PFC_ItemBottomSilkScreenProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCBottomSilkScreenProces.UpdatePFCItemBottomSilkScreenProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCBottomSilkScreenProcess) DeletePFCItemBottomSilkScreenProcess(ctx *gin.Context) {
	var requestParams types.PFC_ItemBottomSilkScreenProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCBottomSilkScreenProces.DeletePFCItemBottomSilkScreenProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
