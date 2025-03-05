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

type PFCKeyManufacturingDetails struct {
	*BaseController
}

var PFCKeyManufacturingDetail = &PFCKeyManufacturingDetails{}

func (c *PFCKeyManufacturingDetails) GetAllPFCKeyManufacturingDetails(ctx *gin.Context) {
	var requestParams *types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCKeyManufacturingDetail.GetAllPFCKeyManufacturingDetails(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCKeyManufacturingDetails) InsertPFCKeyManufacturingDetails(ctx *gin.Context) {
	var requestParams types.PFC_KeyManufacturingDetails
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCKeyManufacturingDetail.InsertNewPFCKeyManufacturingDetails(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCKeyManufacturingDetails) UpdatePFCKeyManufacturingDetails(ctx *gin.Context) {
	var requestParams types.PFC_KeyManufacturingDetails
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCKeyManufacturingDetail.UpdatePFCKeyManufacturingDetails(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCKeyManufacturingDetails) DeletePFCKeyManufacturingDetails(ctx *gin.Context) {
	var requestParams types.PFC_KeyManufacturingDetails
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCKeyManufacturingDetail.DeletePFCKeyManufacturingDetails(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

// //ITEM PFC Key Manufacturing Details
func (c *PFCKeyManufacturingDetails) InsertPFCItemKeyManufacturingDetails(ctx *gin.Context) {
	var requestParams types.PFC_ItemKeyManufacturingDetails
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}
	result, err := services.PFCKeyManufacturingDetail.InsertNewPFCItemKeyManufacturingDetails(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCKeyManufacturingDetails) GetAllPFCItemKeyManufacturingDetails(ctx *gin.Context) {
	var requestParams *types.PFC_KeyManufacturingDetails
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCKeyManufacturingDetail.GetAllPFCItemKeyManufacturingDetails(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCKeyManufacturingDetails) UpdatePFCItemKeyManufacturingDetails(ctx *gin.Context) {
	var requestParams types.PFC_ItemKeyManufacturingDetails
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCKeyManufacturingDetail.UpdatePFCItemKeyManufacturingDetails(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCKeyManufacturingDetails) DeletePFCItemKeyManufacturingDetails(ctx *gin.Context) {
	var requestParams types.PFC_ItemKeyManufacturingDetails
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCKeyManufacturingDetail.DeletePFCItemKeyManufacturingDetails(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
