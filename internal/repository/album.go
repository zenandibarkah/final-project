package repository

import (
	"context"
	"final-project/internal/entity"
)

func (repo *allRepository) GetAllAlbumByArtistID(ctx context.Context, artist_id int64) ([]entity.Album, error) {
	return repo.postgres.GetAllAlbumByArtistID(ctx, artist_id)
}

func (repo *allRepository) GetAllAlbum(ctx context.Context) ([]entity.Album, error) {
	return repo.postgres.GetAllAlbum(ctx)
}

func (repo *allRepository) GetAlbum(ctx context.Context, id int64) (*entity.Album, error) {
	return repo.postgres.GetAlbum(ctx, id)
}

func (repo *allRepository) CreateAlbum(ctx context.Context, album *entity.Album) (int64, error) {
	return repo.postgres.CreateAlbum(ctx, album)
}

func (repo *allRepository) UpdateAlbum(ctx context.Context, album entity.Album) error {
	return repo.postgres.UpdateAlbum(ctx, album)
}

func (repo *allRepository) DeleteAlbum(ctx context.Context, id int64) error {
	return repo.postgres.DeleteAlbum(ctx, id)
}

// CACHE
func (repo *allRepository) GetAlbumCache(ctx context.Context, id int64) (*entity.Album, error) {
	return repo.cache.GetAlbumCache(ctx, id)
}

func (repo *allRepository) SetAlbumCache(ctx context.Context, id int64, album entity.Album) error {
	return repo.cache.SetAlbumCache(ctx, id, album)
}

func (repo *allRepository) DeleteAlbumCache(ctx context.Context, id int64) error {
	return repo.cache.DeleteAlbumCache(ctx, id)
}
