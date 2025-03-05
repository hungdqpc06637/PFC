package controllers

import (
	"log"
	"net/http"
	"web-api/internal/api/services"
	"web-api/internal/pkg/models/response"
	"web-api/internal/pkg/models/types"

	"github.com/gin-gonic/gin"
)

type PFCElasticGoreSpecificationController struct {
	*BaseController
}

var PFCElasticGoreSpecification = &PFCElasticGoreSpecificationController{}

func (c *PFCElasticGoreSpecificationController) GetAllPFCElasticGoreSpecification(ctx *gin.Context) {
	var requestParams *types.PFCModel
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCElasticGoreSpecification.GetAllPFCElasticGoreSpecification(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCElasticGoreSpecificationController) InsertPFCElasticGoreSpecification(ctx *gin.Context) {
	var requestParams types.PFC_ElasticGoreSpecification
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCElasticGoreSpecification.InsertNewPFCElasticGoreSpecification(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCElasticGoreSpecificationController) UpdatePFCElasticGoreSpecification(ctx *gin.Context) {
	var requestParams types.PFC_ElasticGoreSpecification
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCElasticGoreSpecification.UpdatePFCElasticGoreSpecification(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCElasticGoreSpecificationController) DeletePFCElasticGoreSpecification(ctx *gin.Context) {
	var requestParams types.PFC_ElasticGoreSpecification
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCElasticGoreSpecification.DeletePFCElasticGoreSpecification(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

//ITEM PFC Upper Logo Specification

func (c *PFCElasticGoreSpecificationController) InsertPFCItemElasticGoreSpecification(ctx *gin.Context) {
	var requestParams types.PFC_ItemElasticGoreSpecification
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCElasticGoreSpecification.InsertNewPFCItemElasticGoreSpecification(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCElasticGoreSpecificationController) GetAllPFCItemElasticGoreSpecification(ctx *gin.Context) {
	var requestParams *types.PFC_ElasticGoreSpecification
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCElasticGoreSpecification.GetAllPFCItemElasticGoreSpecification(requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCElasticGoreSpecificationController) UpdatePFCItemElasticGoreSpecification(ctx *gin.Context) {
	var requestParams types.PFC_ItemElasticGoreSpecification
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCElasticGoreSpecification.UpdatePFCItemElasticGoreSpecification(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *PFCElasticGoreSpecificationController) DeletePFCItemElasticGoreSpecification(ctx *gin.Context) {
	var requestParams types.PFC_ItemElasticGoreSpecification
	if err := ctx.ShouldBind(&requestParams); err != nil {
		log.Printf("Invalid form data: %v", err)
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	result, err := services.PFCElasticGoreSpecification.DeletePFCItemElasticGoreSpecification(&requestParams)

	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
