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

type PFCSocklinerMoldingProcess struct {
	*BaseController
}

var PFCSocklinerMoldingProces = &PFCSocklinerMoldingProcess{}

func (c *PFCSocklinerMoldingProcess) GetAllPFCSocklinerMoldingProcess(ctx *gin.Context) {
	var requestParams *types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCSocklinerMoldingProces.GetAllPFCSocklinerMoldingProcess(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCSocklinerMoldingProcess) InsertPFCSocklinerMoldingProcess(ctx *gin.Context) {
	var requestParams types.PFC_SocklinerMoldingProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCSocklinerMoldingProces.InsertNewPFCSocklinerMoldingProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCSocklinerMoldingProcess) UpdatePFCSocklinerMoldingProcess(ctx *gin.Context) {
	var requestParams types.PFC_SocklinerMoldingProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCSocklinerMoldingProces.UpdatePFCSocklinerMoldingProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCSocklinerMoldingProcess) DeletePFCSocklinerMoldingProcess(ctx *gin.Context) {
	var requestParams types.PFC_SocklinerMoldingProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCSocklinerMoldingProces.DeletePFCSocklinerMoldingProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

// //ITEM PFC Sockliner Molding Process
func (c *PFCSocklinerMoldingProcess) InsertPFCItemSocklinerMoldingProcess(ctx *gin.Context) {
	var requestParams types.PFC_ItemSocklinerMoldingProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}
	result, err := services.PFCSocklinerMoldingProces.InsertNewPFCItemSocklinerMoldingProcess(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCSocklinerMoldingProcess) GetAllPFCItemSocklinerMoldingProcess(ctx *gin.Context) {
	var requestParams *types.PFC_SocklinerMoldingProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCSocklinerMoldingProces.GetAllPFCItemSocklinerMoldingProcess(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCSocklinerMoldingProcess) UpdatePFCItemSocklinerMoldingProcess(ctx *gin.Context) {
	var requestParams types.PFC_ItemSocklinerMoldingProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCSocklinerMoldingProces.UpdatePFCItemSocklinerMoldingProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCSocklinerMoldingProcess) DeletePFCItemSocklinerMoldingProcess(ctx *gin.Context) {
	var requestParams types.PFC_ItemSocklinerMoldingProcess
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCSocklinerMoldingProces.DeletePFCItemSocklinerMoldingProcess(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
