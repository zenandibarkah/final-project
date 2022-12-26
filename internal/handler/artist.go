package handler

import (
	"final-project/internal/entity"
	"final-project/internal/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (handler *allHandler) GetArtist(context *gin.Context) {
	// Get id from request param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Call the usecase
	artist, err := handler.allUsecase.GetArtist(context, id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", artist)
	context.JSON(http.StatusOK, res)
}

func (handler *allHandler) GetAllArtist(context *gin.Context) {
	// Get all artists from usecase
	artists, err := handler.allUsecase.GetAllArtist(context)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", artists)
	context.JSON(http.StatusOK, res)
}

func (handler *allHandler) CreateArtist(context *gin.Context) {
	var requestBody entity.Artist

	// Get request body from user
	err := context.BindJSON(&requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Call the usecase
	artist, err := handler.allUsecase.CreateArtist(context, &requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", artist)
	context.JSON(http.StatusOK, res)
}

func (handler *allHandler) BatchCreateArtist(context *gin.Context) {
	var requestBody []entity.Artist

	// Get request body from user
	err := context.BindJSON(&requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	artists, err := handler.allUsecase.BatchCreateArtist(context, requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", artists)
	context.JSON(http.StatusOK, res)
}

func (handler *allHandler) UpdateArtist(context *gin.Context) {
	var requestBody entity.Artist

	// Get id from request param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Get request body from user
	err = context.BindJSON(&requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Set id from params
	requestBody.ID = id

	// Call the usecase
	artist, err := handler.allUsecase.UpdateArtist(context, requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", artist)
	context.JSON(http.StatusOK, res)
}

func (handler *allHandler) DeleteArtist(context *gin.Context) {
	// Get id from request param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = handler.allUsecase.DeleteArtist(context, id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
