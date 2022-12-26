package usecase

import (
	"context"
	"final-project/internal/entity"
)

func (usecase *allUsecase) GetAllAlbumByArtistID(ctx context.Context, artist_id int64) ([]entity.Album, error) {
	var albums []entity.Album

	// Get Album from DB
	albums, err := usecase.allRepository.GetAllAlbumByArtistID(ctx, artist_id)
	if err != nil {
		return albums, err
	}

	return albums, nil
}

func (usecase *allUsecase) GetAllAlbum(ctx context.Context) ([]entity.Album, error) {
	var albums []entity.Album

	// Get Album from DB
	albums, err := usecase.allRepository.GetAllAlbum(ctx)
	if err != nil {
		return albums, err
	}

	return albums, nil
}

func (usecase *allUsecase) GetAlbum(ctx context.Context, id int64) (*entity.Album, error) {
	// Get from cache
	album, err := usecase.allRepository.GetAlbumCache(ctx, id)
	if err != nil {
		return album, err
	}

	if album.ID != 0 {
		return album, nil
	}

	// get from DB
	album, err = usecase.allRepository.GetAlbum(ctx, id)
	if err != nil {
		return album, err
	}

	if err = usecase.allRepository.SetAlbumCache(ctx, id, *album); err != nil {
		return album, err
	}

	return album, nil
}

func (usecase *allUsecase) CreateAlbum(ctx context.Context, album *entity.Album) (*entity.Album, error) {
	var NewAlbum *entity.Album

	// Create artist to DB
	id, err := usecase.allRepository.CreateAlbum(ctx, album)
	if err != nil {
		return NewAlbum, err
	}

	// Find new artist from DB
	NewAlbum, err = usecase.allRepository.GetAlbum(ctx, id)
	if err != nil {
		return NewAlbum, err
	}

	// Set to specific cache
	if err = usecase.allRepository.SetAlbumCache(ctx, id, *NewAlbum); err != nil {
		return NewAlbum, err
	}

	return NewAlbum, nil
}

func (usecase *allUsecase) UpdateAlbum(ctx context.Context, album entity.Album) (entity.Album, error) {
	var updateAlbum *entity.Album

	// Update album
	err := usecase.allRepository.UpdateAlbum(ctx, album)
	if err != nil {
		return *updateAlbum, err
	}

	// Find new album
	updateAlbum, err = usecase.allRepository.GetAlbum(ctx, album.ID)
	if err != nil {
		return *updateAlbum, err
	}

	// Set to specific cache
	if err = usecase.allRepository.SetAlbumCache(ctx, album.ID, *updateAlbum); err != nil {
		return *updateAlbum, err
	}

	return *updateAlbum, nil
}

func (usecase *allUsecase) DeleteAlbum(ctx context.Context, id int64) error {
	// Delete from DB
	err := usecase.allRepository.DeleteAlbum(ctx, id)
	if err != nil {
		return err
	}

	// Delete cache
	err = usecase.allRepository.DeleteAlbumCache(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
