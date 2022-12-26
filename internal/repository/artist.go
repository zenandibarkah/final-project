package repository

import (
	"context"
	"final-project/internal/entity"
)

func (repo *allRepository) GetArtist(ctx context.Context, id int64) (*entity.Artist, error) {
	return repo.postgres.GetArtist(ctx, id)
}

func (repo *allRepository) GetAllArtist(ctx context.Context) ([]entity.Artist, error) {
	return repo.postgres.GetAllArtist(ctx)
}

func (repo *allRepository) CreateArtist(ctx context.Context, artist *entity.Artist) (int64, error) {
	return repo.postgres.CreateArtist(ctx, artist)
}

func (repo *allRepository) BatchCreateArtist(ctx context.Context, artists []entity.Artist) ([]int64, error) {
	return repo.postgres.BatchCreateArtist(ctx, artists)
}

func (repo *allRepository) UpdateArtist(ctx context.Context, artist entity.Artist) error {
	return repo.postgres.UpdateArtist(ctx, artist)
}

func (repo *allRepository) DeleteArtist(ctx context.Context, id int64) error {
	return repo.postgres.DeleteArtist(ctx, id)
}

// CACHE
func (repo *allRepository) GetArtistCache(ctx context.Context, id int64) (*entity.Artist, error) {
	return repo.cache.GetArtistCache(ctx, id)
}

func (repo *allRepository) SetArtistCache(ctx context.Context, id int64, artist entity.Artist) error {
	return repo.cache.SetArtistCache(ctx, id, artist)
}

func (repo *allRepository) DeleteArtistCache(ctx context.Context, id int64) error {
	return repo.cache.DeleteArtistCache(ctx, id)
}
