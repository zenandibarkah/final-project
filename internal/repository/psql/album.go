package psql

import (
	"context"
	"final-project/internal/entity"
	"fmt"

	"time"
)

func (repo *postgresConnection) GetAllAlbumByArtistID(ctx context.Context, artist_id int64) ([]entity.Album, error) {
	query := `
		SELECT albums.id, albums.title, albums.price, albums.artist_id, artists.name as artist_name
		FROM albums INNER JOIN artists ON artists.id = albums.artist_id
		WHERE albums.artist_id = $1`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var albums []entity.Album

	rows, err := repo.db.QueryContext(ctx, query, artist_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var album entity.Album

		err := rows.Scan(
			&album.ID,
			&album.Title,
			&album.Price,
			&album.DataArtist.Artist_id,
			&album.DataArtist.Name,
		)

		if err != nil {
			return nil, err
		}

		albums = append(albums, album)
	}
	// fmt.Println(albums)
	return albums, nil
}

func (repo *postgresConnection) GetAllAlbum(ctx context.Context) ([]entity.Album, error) {
	query := `
		SELECT albums.id, albums.title, albums.price, albums.artist_id, artists.name as artist_name
		FROM albums INNER JOIN artists ON artists.id = albums.artist_id`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var albums []entity.Album

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var album entity.Album

		err := rows.Scan(
			&album.ID,
			&album.Title,
			&album.Price,
			&album.DataArtist.Artist_id,
			&album.DataArtist.Name,
		)

		if err != nil {
			return nil, err
		}

		albums = append(albums, album)
	}
	// fmt.Println(albums)
	return albums, nil
}

func (repo *postgresConnection) GetAlbum(ctx context.Context, id int64) (*entity.Album, error) {
	query := `
		SELECT albums.id, albums.title, albums.price, albums.artist_id, artists.name as artist_name
		FROM albums INNER JOIN artists ON artists.id = albums.artist_id
		WHERE albums.id = $1`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var album entity.Album

	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		// Scan data from DB with query
		&album.ID,
		&album.Title,
		&album.Price,
		&album.DataArtist.Artist_id,
		&album.DataArtist.Name,
	)
	if err != nil {
		return nil, err
	}

	return &album, nil
}

func (repo *postgresConnection) CreateAlbum(ctx context.Context, album *entity.Album) (int64, error) {
	query := `
		INSERT INTO public.albums (artist_id, title, price) 
        VALUES ($1, $2, $3)
        RETURNING id`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Run the query insert
	err := repo.db.QueryRowContext(ctx, query, album.DataArtist.Artist_id, album.Title, album.Price).Scan(&album.ID)
	if err != nil {
		return 0, err
	}

	return album.ID, nil
}

func (repo *postgresConnection) UpdateAlbum(ctx context.Context, album entity.Album) error {
	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query update
	query := `UPDATE albums set artist_id=$1, title=$2, price=$3 WHERE id=$4`

	// Run the query
	result, err := repo.db.ExecContext(ctx, query, album.DataArtist.Artist_id, album.Title, album.Price, album.ID)
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

func (repo *postgresConnection) DeleteAlbum(ctx context.Context, id int64) error {
	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query delete
	query := `DELETE from albums WHERE id=$1`

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
