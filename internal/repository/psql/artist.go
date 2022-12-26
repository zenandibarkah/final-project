package psql

import (
	"context"
	"final-project/internal/entity"
	"fmt"
	"log"
	"time"
)

func (repo *postgresConnection) GetArtist(ctx context.Context, id int64) (*entity.Artist, error) {
	query := `
		SELECT id, name
		FROM artists
		where id = $1`

	var artist entity.Artist

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		&artist.ID,
		&artist.Name,
	)

	if err != nil {
		return nil, err
	}

	return &artist, nil
}

func (repo *postgresConnection) GetAllArtist(ctx context.Context) ([]entity.Artist, error) {
	query := `
		SELECT id, name
		FROM artists`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var artists []entity.Artist

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var artist entity.Artist

		err := rows.Scan(
			&artist.ID,
			&artist.Name,
		)

		if err != nil {
			return nil, err
		}

		artists = append(artists, artist)
	}

	return artists, nil
}

func (repo *postgresConnection) CreateArtist(ctx context.Context, artist *entity.Artist) (int64, error) {
	// The query insert
	query := `
        INSERT INTO public.artists (name) 
        VALUES ($1)
        RETURNING id`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Run the query insert
	err := repo.db.QueryRowContext(ctx, query, artist.Name).Scan(&artist.ID)
	if err != nil {
		return 0, err
	}

	return artist.ID, nil
}

func (repo *postgresConnection) BatchCreateArtist(ctx context.Context, artists []entity.Artist) ([]int64, error) {
	var IDArtists []int64

	// Begin transaction
	tx, err := repo.db.Begin()
	if err != nil {
		return IDArtists, nil
	}
	// If any error, the transaction will be rollback
	defer tx.Rollback()

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query insert
	query := `INSERT INTO artists (name) VALUES ($1) RETURNING id`

	// Loop every album
	for _, artist := range artists {
		var id_artist int64

		// Run query insert of every album to database
		err := tx.QueryRowContext(ctx, query, artist.Name).Scan(&id_artist)
		if err != nil {
			log.Printf("error execute insert err: %v", err)
			continue
		}

		// Add the new id to IDs variable
		IDArtists = append(IDArtists, id_artist)
	}

	// Commit the transaction
	err = tx.Commit()
	// If any error
	if err != nil {
		return IDArtists, err
	}

	return IDArtists, nil
}

func (repo *postgresConnection) UpdateArtist(ctx context.Context, artist entity.Artist) error {
	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query update
	query := `UPDATE artists set name=$1 WHERE id=$2`

	// Run the query
	result, err := repo.db.ExecContext(ctx, query, artist.Name, artist.ID)
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

func (repo *postgresConnection) DeleteArtist(ctx context.Context, id int64) error {
	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query delete
	query := `DELETE from artists WHERE id=$1`

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
