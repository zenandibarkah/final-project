package main

import (
	"final-project/internal/config"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Get the environment
	env := os.Getenv("ENV")
	if env == "production" || env == "staging" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		// Get the config from .env file
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	// Initialize gin
	r := gin.Default()

	port := os.Getenv("PORT")

	// Load db config
	db, err := config.OpenDB(os.Getenv("POSTGRES_URL"), true)
	if err != nil {
		log.Fatalln(err)
	}
	defer config.CloseDB(db)

	// Load redis
	cache, err := config.OpenCache(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	// Init clean arch
	repository := config.InitRepository(db, cache)
	usecase := config.InitUsecase(repository.AllRepository)
	handler := config.InitHandler(usecase.AllUsecase)

	// Create the API
	artistRoutes := r.Group("/api/v1/artists")
	{
		artistRoutes.POST("/", handler.AllHandler.CreateArtist)
		artistRoutes.POST("/batch", handler.AllHandler.BatchCreateArtist)
		artistRoutes.GET("/", handler.AllHandler.GetAllArtist)
		artistRoutes.GET("/:id", handler.AllHandler.GetArtist)
		artistRoutes.PUT("/:id", handler.AllHandler.UpdateArtist)
		artistRoutes.DELETE("/:id", handler.AllHandler.DeleteArtist)
	}

	albumRoutes := r.Group("/api/v1/albums")
	{
		albumRoutes.GET("/", handler.AllHandler.GetAllAlbum)
		albumRoutes.GET("/:id", handler.AllHandler.GetAlbum)
		albumRoutes.POST("/", handler.AllHandler.CreateAlbum)
		albumRoutes.PUT("/:id", handler.AllHandler.UpdateAlbum)
		albumRoutes.DELETE("/:id", handler.AllHandler.DeleteAlbum)
	}

	songRoutes := r.Group("/api/v1/songs")
	{
		songRoutes.GET("/", handler.AllHandler.GetAllSong)
		songRoutes.GET("/:id", handler.AllHandler.GetSong)
		songRoutes.POST("/", handler.AllHandler.CreateSong)
		songRoutes.PUT("/:id", handler.AllHandler.UpdateSong)
		songRoutes.DELETE("/:id", handler.AllHandler.DeleteSong)
	}

	// Run the gin gonic in port 5000
	runWithPort := fmt.Sprintf("0.0.0.0:%s", port)
	r.Run(runWithPort)

}
