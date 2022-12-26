package handler

import (
	"final-project/internal/entity"
	"final-project/internal/helper"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

func (handler *allHandler) GetAllSong(context *gin.Context) {
	reqQueryURL := context.Request.URL.RawQuery

	if reqQueryURL != "" {
		album_id_query := context.Query("album_id")
		artist_id_query := context.Query("artist_id")
		if album_id_query != "" && artist_id_query == "" {
			album_id, _ := strconv.ParseInt(album_id_query, 10, 64)

			songs, err := handler.allUsecase.GetAllSongByAlbumID(context, album_id)
			if err != nil {
				res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
				context.AbortWithStatusJSON(http.StatusInternalServerError, res)
				return
			}
			res := helper.BuildResponse(true, "OK!", songs)
			context.JSON(http.StatusOK, res)

		} else if album_id_query == "" && artist_id_query != "" {
			artist_id, _ := strconv.ParseInt(artist_id_query, 10, 64)

			songs, err := handler.allUsecase.GetAllSongByArtistID(context, artist_id)
			if err != nil {
				res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
				context.AbortWithStatusJSON(http.StatusInternalServerError, res)
				return
			}
			res := helper.BuildResponse(true, "OK!", songs)
			context.JSON(http.StatusOK, res)

		} else if album_id_query == "" && artist_id_query == "" {
			res := helper.BuildErrorResponse(
				"Error Query Params",
				"No Query album_id or artist_id",
				helper.EmptyObj{},
			)
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

	} else {
		songs, err := handler.allUsecase.GetAllSong(context)
		if err != nil {
			res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}

		res := helper.BuildResponse(true, "OK!", songs)
		context.JSON(http.StatusOK, res)
	}
}

func (handler *allHandler) GetSong(context *gin.Context) {
	// Get id from request param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Get all albums from usecase
	song, err := handler.allUsecase.GetSong(context, id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", song)
	context.JSON(http.StatusOK, res)
}

func (handler *allHandler) CreateSong(context *gin.Context) {
	var requestBody entity.Song

	// Get request body from user
	err := context.BindJSON(&requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Call the usecase
	song, err := handler.allUsecase.CreateSong(context, &requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", song)
	context.JSON(http.StatusOK, res)
}

func (handler *allHandler) UpdateSong(context *gin.Context) {
	var requestBody entity.Song

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
	song, err := handler.allUsecase.UpdateSong(context, requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", song)
	context.JSON(http.StatusOK, res)
}

func (handler *allHandler) DeleteSong(context *gin.Context) {
	// Get id from request param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = handler.allUsecase.DeleteSong(context, id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
