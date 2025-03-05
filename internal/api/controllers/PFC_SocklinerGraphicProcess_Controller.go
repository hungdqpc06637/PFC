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

type PFCSocklinerGraphicProcess struct {
	*BaseController
}

var PFCSocklinerGraphicProces = &PFCSocklinerGraphicProcess{}

func (c *PFCSocklinerGraphicProcess) GetAllPFCSocklinerGraphicProcess(ctx *gin.Context) {
	var requestParams *types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCSocklinerGraphicProces.GetAllPFCSocklinerGraphicProcess(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCSocklinerGraphicProcess) InsertPFCSocklinerGraphicProcess(ctx *gin.Context) {
	var requestParams types.PFC_SocklinerGraphicProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCSocklinerGraphicProces.InsertNewPFCSocklinerGraphicProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCSocklinerGraphicProcess) UpdatePFCSocklinerGraphicProcess(ctx *gin.Context) {
	var requestParams types.PFC_SocklinerGraphicProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCSocklinerGraphicProces.UpdatePFCSocklinerGraphicProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCSocklinerGraphicProcess) DeletePFCSocklinerGraphicProcess(ctx *gin.Context) {
	var requestParams types.PFC_SocklinerGraphicProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCSocklinerGraphicProces.DeletePFCSocklinerGraphicProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

// //ITEM PFC Sockliner Molding Process
func (c *PFCSocklinerGraphicProcess) InsertPFCItemSocklinerGraphicProcess(ctx *gin.Context) {
	var requestParams types.PFC_ItemSocklinerGraphicProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}
	result, err := services.PFCSocklinerGraphicProces.InsertNewPFCItemSocklinerGraphicProcess(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCSocklinerGraphicProcess) GetAllPFCItemSocklinerGraphicProcess(ctx *gin.Context) {
	var requestParams *types.PFC_SocklinerGraphicProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCSocklinerGraphicProces.GetAllPFCItemSocklinerGraphicProcess(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCSocklinerGraphicProcess) UpdatePFCItemSocklinerGraphicProcess(ctx *gin.Context) {
	var requestParams types.PFC_ItemSocklinerGraphicProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCSocklinerGraphicProces.UpdatePFCItemSocklinerGraphicProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCSocklinerGraphicProcess) DeletePFCItemSocklinerGraphicProcess(ctx *gin.Context) {
	var requestParams types.PFC_ItemSocklinerGraphicProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCSocklinerGraphicProces.DeletePFCItemSocklinerGraphicProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
