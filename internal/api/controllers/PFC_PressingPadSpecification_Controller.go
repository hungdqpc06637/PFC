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

type PFCPressingPadSpecification struct {
	*BaseController
}

var PFCPressingPadSpecificatio = &PFCPressingPadSpecification{}

func (c *PFCPressingPadSpecification) GetAllPFCPressingPadSpecification(ctx *gin.Context) {
	var requestParams *types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCPressingPadSpecificatio.GetAllPFCPressingPadSpecification(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCPressingPadSpecification) InsertPFCPressingPadSpecification(ctx *gin.Context) {
	var requestParams types.PFC_PressingPadSpecification
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCPressingPadSpecificatio.InsertNewPFCPressingPadSpecification(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCPressingPadSpecification) UpdatePFCPressingPadSpecification(ctx *gin.Context) {
	var requestParams types.PFC_PressingPadSpecification
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCPressingPadSpecificatio.UpdatePFCPressingPadSpecification(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCPressingPadSpecification) DeletePFCPressingPadSpecification(ctx *gin.Context) {
	var requestParams types.PFC_PressingPadSpecification
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCPressingPadSpecificatio.DeletePFCPressingPadSpecification(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

// //ITEM PFC Pressing Pad Specification
func (c *PFCPressingPadSpecification) InsertPFCItemPressingPadSpecification(ctx *gin.Context) {
	var requestParams types.PFC_ItemPressingPadSpecification
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}
	result, err := services.PFCPressingPadSpecificatio.InsertNewPFCItemPressingPadSpecification(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *PFCPressingPadSpecification) GetAllPFCItemPressingPadSpecification(ctx *gin.Context) {
	var requestParams *types.PFC_PressingPadSpecification
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCPressingPadSpecificatio.GetAllPFCItemPressingPadSpecification(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCPressingPadSpecification) UpdatePFCItemPressingPadSpecification(ctx *gin.Context) {
	var requestParams types.PFC_ItemPressingPadSpecification
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCPressingPadSpecificatio.UpdatePFCItemPressingPadSpecification(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCPressingPadSpecification) DeletePFCItemPressingPadSpecification(ctx *gin.Context) {
	var requestParams types.PFC_ItemPressingPadSpecification
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCPressingPadSpecificatio.DeletePFCItemPressingPadSpecification(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
