package cache

import (
	"context"
	"final-project/internal/entity"

	"github.com/go-redis/redis/v8"
)

// ARTIST
type ClientRedis interface {
	// ARTIST
	GetArtistCache(ctx context.Context, id int64) (*entity.Artist, error)
	SetArtistCache(ctx context.Context, id int64, artist entity.Artist) error
	DeleteArtistCache(ctx context.Context, id int64) error

	// ALBUM
	GetAlbumCache(ctx context.Context, id int64) (*entity.Album, error)
	SetAlbumCache(ctx context.Context, id int64, album entity.Album) error
	DeleteAlbumCache(ctx context.Context, id int64) error

	// SONG
	GetSongCache(ctx context.Context, id int64) (*entity.Song, error)
	SetSongCache(ctx context.Context, id int64, song entity.Song) error
	DeleteSongCache(ctx context.Context, id int64) error
}

type clientConnection struct {
	client *redis.Client
}

func NewAClientRedis(cache *redis.Client) ClientRedis {
	return &clientConnection{
		client: cache,
	}
}
