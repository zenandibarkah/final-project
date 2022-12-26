package repository

import (
	"context"
	"database/sql"
	"final-project/internal/entity"
	"final-project/internal/repository/cache"
	"final-project/internal/repository/psql"

	"github.com/go-redis/redis/v8"
)

type AllRepository interface {
	// ARTIST
	GetArtist(ctx context.Context, id int64) (*entity.Artist, error)
	GetAllArtist(ctx context.Context) ([]entity.Artist, error)
	CreateArtist(ctx context.Context, artist *entity.Artist) (int64, error)
	BatchCreateArtist(ctx context.Context, artists []entity.Artist) ([]int64, error)
	UpdateArtist(ctx context.Context, artist entity.Artist) error
	DeleteArtist(ctx context.Context, id int64) error

	GetArtistCache(ctx context.Context, id int64) (*entity.Artist, error)
	SetArtistCache(ctx context.Context, id int64, artist entity.Artist) error
	DeleteArtistCache(ctx context.Context, id int64) error

	// ALBUM
	GetAllAlbumByArtistID(ctx context.Context, artist_id int64) ([]entity.Album, error)
	GetAllAlbum(ctx context.Context) ([]entity.Album, error)
	GetAlbum(ctx context.Context, id int64) (*entity.Album, error)
	CreateAlbum(ctx context.Context, album *entity.Album) (int64, error)
	UpdateAlbum(ctx context.Context, album entity.Album) error
	DeleteAlbum(ctx context.Context, id int64) error

	GetAlbumCache(ctx context.Context, id int64) (*entity.Album, error)
	SetAlbumCache(ctx context.Context, id int64, album entity.Album) error
	DeleteAlbumCache(ctx context.Context, id int64) error

	// SONG
	GetAllSong(ctx context.Context) ([]entity.Song, error)
	GetAllSongByAlbumID(ctx context.Context, album_id int64) ([]entity.Song, error)
	GetAllSongByArtistID(ctx context.Context, artist_id int64) ([]entity.Song, error)
	GetSong(ctx context.Context, id int64) (*entity.Song, error)
	CreateSong(ctx context.Context, song *entity.Song) (int64, error)
	UpdateSong(ctx context.Context, song entity.Song) error
	DeleteSong(ctx context.Context, id int64) error

	GetSongCache(ctx context.Context, id int64) (*entity.Song, error)
	SetSongCache(ctx context.Context, id int64, song entity.Song) error
	DeleteSongCache(ctx context.Context, id int64) error
}

type allRepository struct {
	postgres psql.AllPostgres
	cache    cache.ClientRedis
}

func NewRepository(db *sql.DB, client *redis.Client) AllRepository {
	return &allRepository{
		postgres: psql.NewAllPostgres(db),
		cache:    cache.NewAClientRedis(client),
	}
}
