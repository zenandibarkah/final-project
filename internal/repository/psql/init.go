package psql

import (
	"context"
	"database/sql"
	"final-project/internal/entity"
)

// ARTIST
type AllPostgres interface {
	// ARTIST
	GetArtist(ctx context.Context, id int64) (*entity.Artist, error)
	GetAllArtist(ctx context.Context) ([]entity.Artist, error)
	CreateArtist(ctx context.Context, artist *entity.Artist) (int64, error)
	BatchCreateArtist(ctx context.Context, artists []entity.Artist) ([]int64, error)
	UpdateArtist(ctx context.Context, artist entity.Artist) error
	DeleteArtist(ctx context.Context, id int64) error

	// ALBUM
	GetAllAlbumByArtistID(ctx context.Context, artist_id int64) ([]entity.Album, error)
	GetAllAlbum(ctx context.Context) ([]entity.Album, error)
	GetAlbum(ctx context.Context, id int64) (*entity.Album, error)
	CreateAlbum(ctx context.Context, album *entity.Album) (int64, error)
	UpdateAlbum(ctx context.Context, album entity.Album) error
	DeleteAlbum(ctx context.Context, id int64) error

	// SONG
	GetAllSong(ctx context.Context) ([]entity.Song, error)
	GetAllSongByAlbumID(ctx context.Context, album_id int64) ([]entity.Song, error)
	GetAllSongByArtistID(ctx context.Context, artist_id int64) ([]entity.Song, error)
	GetSong(ctx context.Context, id int64) (*entity.Song, error)
	CreateSong(ctx context.Context, song *entity.Song) (int64, error)
	UpdateSong(ctx context.Context, song entity.Song) error
	DeleteSong(ctx context.Context, id int64) error
}

type postgresConnection struct {
	db *sql.DB
}

func NewAllPostgres(db *sql.DB) AllPostgres {
	return &postgresConnection{
		db: db,
	}
}
