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

type PFCOutsideConveyorProcess struct {
	*BaseController
}

var PFCOutsideConveyorProces = &PFCOutsideConveyorProcess{}

func (c *PFCOutsideConveyorProcess) GetAllPFCOutsideConveyorProcess(ctx *gin.Context) {
	var requestParams *types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCOutsideConveyorProces.GetAllPFCOutsideConveyorProcess(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCOutsideConveyorProcess) InsertPFCOutsideConveyorProcess(ctx *gin.Context) {
	var requestParams types.PFC_OutsideConveyorProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCOutsideConveyorProces.InsertNewPFCOutsideConveyorProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCOutsideConveyorProcess) UpdatePFCOutsideConveyorProcess(ctx *gin.Context) {
	var requestParams types.PFC_OutsideConveyorProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCOutsideConveyorProces.UpdatePFCOutsideConveyorProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCOutsideConveyorProcess) DeletePFCOutsideConveyorProcess(ctx *gin.Context) {
	var requestParams types.PFC_OutsideConveyorProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCOutsideConveyorProces.DeletePFCOutsideConveyorProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

// //ITEM PFC Outside Conveyor Process
func (c *PFCOutsideConveyorProcess) InsertPFCItemOutsideConveyorProcess(ctx *gin.Context) {
	var requestParams types.PFC_ItemOutsideConveyorProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}
	result, err := services.PFCOutsideConveyorProces.InsertNewPFCItemOutsideConveyorProcess(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCOutsideConveyorProcess) GetAllPFCItemOutsideConveyorProcess(ctx *gin.Context) {
	var requestParams *types.PFC_OutsideConveyorProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCOutsideConveyorProces.GetAllPFCItemOutsideConveyorProcess(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCOutsideConveyorProcess) UpdatePFCItemOutsideConveyorProcess(ctx *gin.Context) {
	var requestParams types.PFC_ItemOutsideConveyorProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCOutsideConveyorProces.UpdatePFCItemOutsideConveyorProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCOutsideConveyorProcess) DeletePFCItemOutsideConveyorProcess(ctx *gin.Context) {
	var requestParams types.PFC_ItemOutsideConveyorProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCOutsideConveyorProces.DeletePFCItemOutsideConveyorProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
