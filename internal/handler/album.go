package handler

import (
	"final-project/internal/entity"
	"final-project/internal/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (handler *allHandler) GetAllAlbum(context *gin.Context) {
	reqQueryURL := context.Request.URL.RawQuery

	if reqQueryURL != "" {
		artist_id_query := context.Query("artist_id")

		if artist_id_query == "" {
			// if err != nil {
			res := helper.BuildErrorResponse(
				"Error Query Params",
				"No Query artist_id",
				helper.EmptyObj{},
			)
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
			// }
		} else if artist_id_query != "" {
			artist_id, err := strconv.ParseInt(artist_id_query, 10, 64)
			if err != nil {
				res := helper.BuildErrorResponse("No query param artist_id was found", err.Error(), helper.EmptyObj{})
				context.AbortWithStatusJSON(http.StatusBadRequest, res)
				return
			}
			// Get all albums from usecase
			albums, err := handler.allUsecase.GetAllAlbumByArtistID(context, artist_id)
			if err != nil {
				res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
				context.AbortWithStatusJSON(http.StatusInternalServerError, res)
				return
			}
			res := helper.BuildResponse(true, "OK!", albums)
			context.JSON(http.StatusOK, res)
		}

	} else {
		albums, err := handler.allUsecase.GetAllAlbum(context)
		if err != nil {
			res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}

		res := helper.BuildResponse(true, "OK!", albums)
		context.JSON(http.StatusOK, res)
	}
}

func (handler *allHandler) GetAlbum(context *gin.Context) {
	// Get id from request param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Get all albums from usecase
	album, err := handler.allUsecase.GetAlbum(context, id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", album)
	context.JSON(http.StatusOK, res)
}

func (handler *allHandler) CreateAlbum(context *gin.Context) {
	var requestBody entity.Album

	// Get request body from user
	err := context.BindJSON(&requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Call the usecase
	album, err := handler.allUsecase.CreateAlbum(context, &requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", album)
	context.JSON(http.StatusOK, res)
}

func (handler *allHandler) UpdateAlbum(context *gin.Context) {
	var requestBody entity.Album

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
	album, err := handler.allUsecase.UpdateAlbum(context, requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", album)
	context.JSON(http.StatusOK, res)
}

func (handler *allHandler) DeleteAlbum(context *gin.Context) {
	// Get id from request param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = handler.allUsecase.DeleteAlbum(context, id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
