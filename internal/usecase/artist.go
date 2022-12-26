package usecase

import (
	"context"
	"final-project/internal/entity"
)

func (usecase *allUsecase) GetArtist(ctx context.Context, id int64) (*entity.Artist, error) {
	// Get from cache
	artist, err := usecase.allRepository.GetArtistCache(ctx, id)
	if err != nil {
		return artist, err
	}

	if artist.ID != 0 {
		return artist, nil
	}

	// Get from db
	artist, err = usecase.allRepository.GetArtist(ctx, id)
	if err != nil {
		return artist, err
	}

	// Set to cache
	if err = usecase.allRepository.SetArtistCache(ctx, id, *artist); err != nil {
		return artist, err
	}

	return artist, nil
}

func (usecase *allUsecase) GetAllArtist(ctx context.Context) ([]entity.Artist, error) {
	var artists []entity.Artist

	// Get from db
	artists, err := usecase.allRepository.GetAllArtist(ctx)
	if err != nil {
		return artists, err
	}

	return artists, nil
}

func (usecase *allUsecase) CreateArtist(ctx context.Context, artist *entity.Artist) (*entity.Artist, error) {
	var NewArtist *entity.Artist

	// Create artist to DB
	id, err := usecase.allRepository.CreateArtist(ctx, artist)
	if err != nil {
		return NewArtist, err
	}

	// Find new artist from DB
	NewArtist, err = usecase.allRepository.GetArtist(ctx, id)
	if err != nil {
		return NewArtist, err
	}

	// Set to specific cache
	if err = usecase.allRepository.SetArtistCache(ctx, id, *NewArtist); err != nil {
		return NewArtist, err
	}

	return NewArtist, nil
}

func (usecase *allUsecase) BatchCreateArtist(ctx context.Context, artists []entity.Artist) ([]entity.Artist, error) {
	var NewArtists []entity.Artist

	// Batch create and get the new id
	ids, err := usecase.allRepository.BatchCreateArtist(ctx, artists)
	if err != nil {
		return NewArtists, err
	}

	// Get detail artist by ids
	for _, id := range ids {
		// Get from db
		artist, err := usecase.allRepository.GetArtist(ctx, id)
		if err != nil {
			return NewArtists, err
		}

		// Set to specific cache
		if err = usecase.allRepository.SetArtistCache(ctx, id, *artist); err != nil {
			return NewArtists, err
		}

		NewArtists = append(NewArtists, *artist)
	}

	return NewArtists, nil
}

func (usecase *allUsecase) UpdateArtist(ctx context.Context, artist entity.Artist) (entity.Artist, error) {
	var updateArtist *entity.Artist

	// Update album
	err := usecase.allRepository.UpdateArtist(ctx, artist)
	if err != nil {
		return *updateArtist, err
	}

	// Find new album
	updateArtist, err = usecase.allRepository.GetArtist(ctx, artist.ID)
	if err != nil {
		return *updateArtist, err
	}

	// Set to specific cache
	if err = usecase.allRepository.SetArtistCache(ctx, artist.ID, *updateArtist); err != nil {
		return *updateArtist, err
	}

	return *updateArtist, nil
}

func (usecase *allUsecase) DeleteArtist(ctx context.Context, id int64) error {
	// Delete from db
	if err := usecase.allRepository.DeleteArtist(ctx, id); err != nil {
		return err
	}

	// Delete from cache
	if err := usecase.allRepository.DeleteArtistCache(ctx, id); err != nil {
		return err
	}

	return nil
}
