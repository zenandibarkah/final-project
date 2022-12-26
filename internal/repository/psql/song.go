package psql

import (
	"context"
	"final-project/internal/entity"

	"fmt"

	"time"
)

func (repo *postgresConnection) GetAllSong(ctx context.Context) ([]entity.Song, error) {
	query := `
		SELECT songs.id, songs.title, COALESCE(songs.lyric, ''), songs.album_id, albums.title,
		artists.id, artists.name
		FROM songs INNER JOIN albums ON albums.id=songs.album_id
		INNER JOIN artists ON artists.id=albums.artist_id 
		ORDER BY songs.id ASC`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var songs []entity.Song

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var song entity.Song

		err := rows.Scan(
			&song.ID,
			&song.Title,
			&song.Lyric,
			&song.DataAlbum.Album_id,
			&song.DataAlbum.Title,
			&song.DataAlbum.DataArtist.Artist_id,
			&song.DataAlbum.DataArtist.Name,
		)

		if err != nil {
			return nil, err
		}

		songs = append(songs, song)
	}

	return songs, nil
}

func (repo *postgresConnection) GetAllSongByAlbumID(ctx context.Context, album_id int64) ([]entity.Song, error) {
	query := `
		SELECT songs.id, songs.title, COALESCE(songs.lyric, ''), songs.album_id, albums.title,
		artists.id, artists.name
		FROM songs INNER JOIN albums ON albums.id=songs.album_id
		INNER JOIN artists ON artists.id=albums.artist_id 
		WHERE albums.id = $1
		ORDER BY songs.id ASC`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var songs []entity.Song

	rows, err := repo.db.QueryContext(ctx, query, album_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var song entity.Song

		err := rows.Scan(
			&song.ID,
			&song.Title,
			&song.Lyric,
			&song.DataAlbum.Album_id,
			&song.DataAlbum.Title,
			&song.DataAlbum.DataArtist.Artist_id,
			&song.DataAlbum.DataArtist.Name,
		)

		if err != nil {
			return nil, err
		}

		songs = append(songs, song)
	}

	return songs, nil
}

func (repo *postgresConnection) GetAllSongByArtistID(ctx context.Context, artist_id int64) ([]entity.Song, error) {
	query := `
		SELECT songs.id, songs.title, COALESCE(songs.lyric, ''), songs.album_id, albums.title,
		artists.id, artists.name
		FROM songs INNER JOIN albums ON albums.id=songs.album_id
		INNER JOIN artists ON artists.id=albums.artist_id 
		WHERE artists.id = $1
		ORDER BY songs.id ASC`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var songs []entity.Song

	rows, err := repo.db.QueryContext(ctx, query, artist_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var song entity.Song

		err := rows.Scan(
			&song.ID,
			&song.Title,
			&song.Lyric,
			&song.DataAlbum.Album_id,
			&song.DataAlbum.Title,
			&song.DataAlbum.DataArtist.Artist_id,
			&song.DataAlbum.DataArtist.Name,
		)

		if err != nil {
			return nil, err
		}

		songs = append(songs, song)
	}

	return songs, nil
}

func (repo *postgresConnection) GetSong(ctx context.Context, id int64) (*entity.Song, error) {
	query := `
		SELECT songs.id, songs.title, COALESCE(songs.lyric, ''), songs.album_id, albums.title,
		artists.id, artists.name
		FROM songs INNER JOIN albums ON albums.id=songs.album_id
		INNER JOIN artists ON artists.id=albums.artist_id 
		WHERE songs.id = $1
		ORDER BY songs.id ASC`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var song entity.Song

	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		// Scan data from DB with query
		&song.ID,
		&song.Title,
		&song.Lyric,
		&song.DataAlbum.Album_id,
		&song.DataAlbum.Title,
		&song.DataAlbum.DataArtist.Artist_id,
		&song.DataAlbum.DataArtist.Name,
	)
	if err != nil {
		return nil, err
	}

	return &song, nil
}

func (repo *postgresConnection) CreateSong(ctx context.Context, song *entity.Song) (int64, error) {
	query := `
		INSERT INTO public.songs (album_id, title, lyric) 
        VALUES ($1, $2, $3)
        RETURNING id`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Run the query insert
	err := repo.db.QueryRowContext(ctx, query, song.DataAlbum.Album_id, song.Title, song.Lyric).Scan(&song.ID)
	if err != nil {
		return 0, err
	}

	return song.ID, nil
}

func (repo *postgresConnection) UpdateSong(ctx context.Context, song entity.Song) error {
	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query update
	query := `UPDATE songs set album_id=$1, title=$2, lyric=$3 WHERE id=$4`

	// Run the query
	result, err := repo.db.ExecContext(ctx, query, song.DataAlbum.Album_id, song.Title, song.Lyric, song.ID)
	if err != nil {
		return err
	}

	// Get how many data has been updated
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Affected update : %d", rows)
	return nil
}

func (repo *postgresConnection) DeleteSong(ctx context.Context, id int64) error {
	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query delete
	query := `DELETE from songs WHERE id=$1`

	// Run the delete query
	result, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	// Get how many data has been deleted
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Affected delete : %d", rows)
	return nil
}
