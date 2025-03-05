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

type PFCReinforcementPlacement struct {
	*BaseController
}

var PFCReinforcementPlacemen = &PFCReinforcementPlacement{}

func (c *PFCReinforcementPlacement) GetAllPFCReinforcementPlacement(ctx *gin.Context) {
	var requestParams *types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCReinforcementPlacemen.GetAllPFCReinforcementPlacement(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCReinforcementPlacement) InsertPFCReinforcementPlacement(ctx *gin.Context) {
	var requestParams types.PFC_ReinforcementPlacement
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCReinforcementPlacemen.InsertNewPFCReinforcementPlacement(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCReinforcementPlacement) UpdatePFCReinforcementPlacement(ctx *gin.Context) {
	var requestParams types.PFC_ReinforcementPlacement
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCReinforcementPlacemen.UpdatePFCReinforcementPlacement(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCReinforcementPlacement) DeletePFCReinforcementPlacement(ctx *gin.Context) {
	var requestParams types.PFC_ReinforcementPlacement
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCReinforcementPlacemen.DeletePFCReinforcementPlacement(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

// //ITEM PFC Skiving Instructions
func (c *PFCReinforcementPlacement) InsertPFCItemReinforcementPlacement(ctx *gin.Context) {
	var requestParams types.PFC_ItemReinforcementPlacement
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}
	result, err := services.PFCReinforcementPlacemen.InsertNewPFCItemReinforcementPlacement(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCReinforcementPlacement) GetAllPFCItemReinforcementPlacement(ctx *gin.Context) {
	var requestParams *types.PFC_ReinforcementPlacement
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCReinforcementPlacemen.GetAllPFCItemReinforcementPlacement(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCReinforcementPlacement) UpdatePFCItemReinforcementPlacement(ctx *gin.Context) {
	var requestParams types.PFC_ItemReinforcementPlacement
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCReinforcementPlacemen.UpdatePFCItemReinforcementPlacement(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCReinforcementPlacement) DeletePFCItemReinforcementPlacement(ctx *gin.Context) {
	var requestParams types.PFC_ItemReinforcementPlacement
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCReinforcementPlacemen.DeletePFCItemReinforcementPlacement(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}





