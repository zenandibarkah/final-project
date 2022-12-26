package config

import (
	"database/sql"
	repository "final-project/internal/repository"

	"github.com/go-redis/redis/v8"
)

type Repository struct {
	AllRepository repository.AllRepository
}

func InitRepository(db *sql.DB, client *redis.Client) Repository {
	return Repository{
		AllRepository: repository.NewRepository(db, client),
	}
}
