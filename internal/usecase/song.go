package usecase

import (
	"context"
	"final-project/internal/entity"
)

func (usecase *allUsecase) GetAllSong(ctx context.Context) ([]entity.Song, error) {
	var songs []entity.Song

	// Get Songs from DB
	songs, err := usecase.allRepository.GetAllSong(ctx)
	if err != nil {
		return songs, err
	}

	return songs, nil
}

func (usecase *allUsecase) GetAllSongByAlbumID(ctx context.Context, album_id int64) ([]entity.Song, error) {
	var songs []entity.Song

	// Get All song by album_id from DB
	songs, err := usecase.allRepository.GetAllSongByAlbumID(ctx, album_id)
	if err != nil {
		return songs, err
	}

	return songs, nil
}

func (usecase *allUsecase) GetAllSongByArtistID(ctx context.Context, artist_id int64) ([]entity.Song, error) {
	var songs []entity.Song

	// Get All song by artist_id from DB
	songs, err := usecase.allRepository.GetAllSongByArtistID(ctx, artist_id)
	if err != nil {
		return songs, err
	}

	return songs, nil
}

func (usecase *allUsecase) GetSong(ctx context.Context, id int64) (*entity.Song, error) {
	// Get from cache
	song, err := usecase.allRepository.GetSongCache(ctx, id)
	if err != nil {
		return song, err
	}

	if song.ID != 0 {
		return song, nil
	}

	// get from DB
	song, err = usecase.allRepository.GetSong(ctx, id)
	if err != nil {
		return song, err
	}

	if err = usecase.allRepository.SetSongCache(ctx, id, *song); err != nil {
		return song, err
	}

	return song, nil
}

func (usecase *allUsecase) CreateSong(ctx context.Context, song *entity.Song) (*entity.Song, error) {
	var NewSong *entity.Song

	// Create artist to DB
	id, err := usecase.allRepository.CreateSong(ctx, song)
	if err != nil {
		return NewSong, err
	}

	// Find new artist from DB
	NewSong, err = usecase.allRepository.GetSong(ctx, id)
	if err != nil {
		return NewSong, err
	}

	// Set to specific cache
	if err = usecase.allRepository.SetSongCache(ctx, id, *NewSong); err != nil {
		return NewSong, err
	}

	return NewSong, nil
}

func (usecase *allUsecase) UpdateSong(ctx context.Context, song entity.Song) (entity.Song, error) {
	var updateSong *entity.Song

	// Update album
	err := usecase.allRepository.UpdateSong(ctx, song)
	if err != nil {
		return *updateSong, err
	}

	// Find new album
	updateSong, err = usecase.allRepository.GetSong(ctx, song.ID)
	if err != nil {
		return *updateSong, err
	}

	// Set to specific cache
	if err = usecase.allRepository.SetSongCache(ctx, song.ID, *updateSong); err != nil {
		return *updateSong, err
	}

	return *updateSong, nil
}

func (usecase *allUsecase) DeleteSong(ctx context.Context, id int64) error {
	// Delete from DB
	err := usecase.allRepository.DeleteSong(ctx, id)
	if err != nil {
		return err
	}

	// Delete cache
	err = usecase.allRepository.DeleteSongCache(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
