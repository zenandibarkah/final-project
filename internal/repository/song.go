package repository

import (
	"context"
	"final-project/internal/entity"
)

// DB
func (repo *allRepository) GetAllSong(ctx context.Context) ([]entity.Song, error) {
	return repo.postgres.GetAllSong(ctx)
}

func (repo *allRepository) GetAllSongByAlbumID(ctx context.Context, album_id int64) ([]entity.Song, error) {
	return repo.postgres.GetAllSongByAlbumID(ctx, album_id)
}

func (repo *allRepository) GetAllSongByArtistID(ctx context.Context, artist_id int64) ([]entity.Song, error) {
	return repo.postgres.GetAllSongByArtistID(ctx, artist_id)
}

func (repo *allRepository) GetSong(ctx context.Context, id int64) (*entity.Song, error) {
	return repo.postgres.GetSong(ctx, id)
}

func (repo *allRepository) CreateSong(ctx context.Context, song *entity.Song) (int64, error) {
	return repo.postgres.CreateSong(ctx, song)
}

func (repo *allRepository) UpdateSong(ctx context.Context, song entity.Song) error {
	return repo.postgres.UpdateSong(ctx, song)
}

func (repo *allRepository) DeleteSong(ctx context.Context, id int64) error {
	return repo.postgres.DeleteSong(ctx, id)
}

// CACHE
func (repo *allRepository) GetSongCache(ctx context.Context, id int64) (*entity.Song, error) {
	return repo.cache.GetSongCache(ctx, id)
}

func (repo *allRepository) SetSongCache(ctx context.Context, id int64, song entity.Song) error {
	return repo.cache.SetSongCache(ctx, id, song)
}

func (repo *allRepository) DeleteSongCache(ctx context.Context, id int64) error {
	return repo.cache.DeleteSongCache(ctx, id)
}
