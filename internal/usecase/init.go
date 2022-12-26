package usecase

import (
	"context"
	"final-project/internal/entity"
	repository "final-project/internal/repository"
)

// ARTIST
type AllUsecase interface {
	// ARTIST
	GetArtist(ctx context.Context, id int64) (*entity.Artist, error)
	GetAllArtist(ctx context.Context) ([]entity.Artist, error)
	CreateArtist(ctx context.Context, artist *entity.Artist) (*entity.Artist, error)
	BatchCreateArtist(ctx context.Context, artists []entity.Artist) ([]entity.Artist, error)
	UpdateArtist(ctx context.Context, artist entity.Artist) (entity.Artist, error)
	DeleteArtist(ctx context.Context, id int64) error

	// ALBUM
	GetAllAlbumByArtistID(ctx context.Context, artist_id int64) ([]entity.Album, error)
	GetAllAlbum(ctx context.Context) ([]entity.Album, error)
	GetAlbum(ctx context.Context, id int64) (*entity.Album, error)
	CreateAlbum(ctx context.Context, album *entity.Album) (*entity.Album, error)
	UpdateAlbum(ctx context.Context, album entity.Album) (entity.Album, error)
	DeleteAlbum(ctx context.Context, id int64) error

	// SONG
	GetAllSong(ctx context.Context) ([]entity.Song, error)
	GetAllSongByAlbumID(ctx context.Context, album_id int64) ([]entity.Song, error)
	GetAllSongByArtistID(ctx context.Context, artist_id int64) ([]entity.Song, error)
	GetSong(ctx context.Context, id int64) (*entity.Song, error)
	CreateSong(ctx context.Context, song *entity.Song) (*entity.Song, error)
	UpdateSong(ctx context.Context, song entity.Song) (entity.Song, error)
	DeleteSong(ctx context.Context, id int64) error
}

type allUsecase struct {
	allRepository repository.AllRepository
}

func NewUsecase(allRepository repository.AllRepository) AllUsecase {
	return &allUsecase{
		allRepository: allRepository,
	}
}
