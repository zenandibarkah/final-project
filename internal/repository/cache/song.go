package cache

import (
	"context"
	"encoding/json"
	"final-project/internal/entity"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func (repo *clientConnection) GetSongCache(ctx context.Context, id int64) (*entity.Song, error) {
	var song entity.Song

	key := fmt.Sprintf(songDetailKey, id)

	songString, err := repo.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return &song, nil
	}
	if err != nil {
		return &song, err
	}

	err = json.Unmarshal([]byte(songString), &song)
	if err != nil {
		return &song, err
	}

	return &song, nil
}

func (repo *clientConnection) SetSongCache(ctx context.Context, id int64, song entity.Song) error {
	key := fmt.Sprintf(songDetailKey, id)

	songString, err := json.Marshal(song)
	if err != nil {
		return err
	}

	if err := repo.client.Set(ctx, key, songString, expiration).Err(); err != nil {
		return nil
	}

	return nil
}

func (repo *clientConnection) DeleteSongCache(ctx context.Context, id int64) error {
	key := fmt.Sprintf(songDetailKey, id)

	if err := repo.client.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}
