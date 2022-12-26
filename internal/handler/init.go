package handler

import (
	usecase "final-project/internal/usecase"

	"github.com/gin-gonic/gin"
)

// ARTIST
type AllHandler interface {
	// ARTIST
	GetArtist(context *gin.Context)
	GetAllArtist(context *gin.Context)
	CreateArtist(context *gin.Context)
	BatchCreateArtist(context *gin.Context)
	UpdateArtist(context *gin.Context)
	DeleteArtist(context *gin.Context)

	// ALBUM
	GetAllAlbum(context *gin.Context)
	GetAlbum(context *gin.Context)
	CreateAlbum(context *gin.Context)
	UpdateAlbum(context *gin.Context)
	DeleteAlbum(context *gin.Context)

	// SONG
	GetAllSong(context *gin.Context)
	GetSong(context *gin.Context)
	CreateSong(context *gin.Context)
	UpdateSong(context *gin.Context)
	DeleteSong(context *gin.Context)
}

type allHandler struct {
	allUsecase usecase.AllUsecase
}

func NewHandler(allUsecase usecase.AllUsecase) AllHandler {
	return &allHandler{
		allUsecase: allUsecase,
	}
}
