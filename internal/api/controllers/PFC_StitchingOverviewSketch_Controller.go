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

type PFCStitchingOverviewSketch struct {
	*BaseController
}

var PFCStitchingOverviewSketc = &PFCStitchingOverviewSketch{}

func (c *PFCStitchingOverviewSketch) GetAllPFCStitchingOverviewSketch(ctx *gin.Context) {
	var requestParams *types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCStitchingOverviewSketc.GetAllPFCStitchingOverviewSketch(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCStitchingOverviewSketch) InsertPFCStitchingOverviewSketch(ctx *gin.Context) {
	var requestParams types.PFC_StitchingOverviewSketch
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCStitchingOverviewSketc.InsertNewPFCStitchingOverviewSketch(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCStitchingOverviewSketch) UpdatePFCStitchingOverviewSketch(ctx *gin.Context) {
	var requestParams types.PFC_StitchingOverviewSketch
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCStitchingOverviewSketc.UpdatePFCStitchingOverviewSketch(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCStitchingOverviewSketch) DeletePFCStitchingOverviewSketch(ctx *gin.Context) {
	var requestParams types.PFC_StitchingOverviewSketch
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCStitchingOverviewSketc.DeletePFCStitchingOverviewSketch(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

// ITEM PFC StitchingOverviewSketch
func (c *PFCStitchingOverviewSketch) InsertPFCItemStitchingOverviewSketch(ctx *gin.Context) {
	var requestParams types.PFC_ItemStitchingOverviewSketch
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}
	result, err := services.PFCStitchingOverviewSketc.InsertNewPFCItemStitchingOverviewSketch(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCStitchingOverviewSketch) GetAllPFCItemStitchingOverviewSketch(ctx *gin.Context) {
	var requestParams *types.PFC_StitchingOverviewSketch
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCStitchingOverviewSketc.GetAllPFCItemStitchingOverviewSketch(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCStitchingOverviewSketch) UpdatePFCItemStitchingOverviewSketch(ctx *gin.Context) {
	var requestParams types.PFC_ItemStitchingOverviewSketch
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCStitchingOverviewSketc.UpdatePFCItemStitchingOverviewSketch(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCStitchingOverviewSketch) DeletePFCItemStitchingOverviewSketch(ctx *gin.Context) {
	var requestParams types.PFC_ItemStitchingOverviewSketch
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCStitchingOverviewSketc.DeletePFCItemStitchingOverviewSketch(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
