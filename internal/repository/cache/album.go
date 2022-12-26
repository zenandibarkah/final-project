package cache

import (
	"context"
	"encoding/json"
	"final-project/internal/entity"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func (repo *clientConnection) GetAlbumCache(ctx context.Context, id int64) (*entity.Album, error) {
	var album entity.Album

	key := fmt.Sprintf(albumDetailKey, id)

	albumString, err := repo.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return &album, nil
	}
	if err != nil {
		return &album, err
	}

	err = json.Unmarshal([]byte(albumString), &album)
	if err != nil {
		return &album, err
	}

	return &album, nil
}

func (repo *clientConnection) SetAlbumCache(ctx context.Context, id int64, album entity.Album) error {
	key := fmt.Sprintf(albumDetailKey, id)

	albumString, err := json.Marshal(album)
	if err != nil {
		return err
	}

	if err := repo.client.Set(ctx, key, albumString, expiration).Err(); err != nil {
		return nil
	}

	return nil
}

func (repo *clientConnection) DeleteAlbumCache(ctx context.Context, id int64) error {
	key := fmt.Sprintf(albumDetailKey, id)

	if err := repo.client.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}
