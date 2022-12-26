package cache

import (
	"context"
	"encoding/json"
	"final-project/internal/entity"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// Get Specific artist cache
func (repo *clientConnection) GetArtistCache(ctx context.Context, id int64) (*entity.Artist, error) {
	var artist entity.Artist

	key := fmt.Sprintf(artistDetailKey, id)

	artistString, err := repo.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return &artist, nil
	}
	if err != nil {
		return &artist, err
	}

	err = json.Unmarshal([]byte(artistString), &artist)
	if err != nil {
		return &artist, err
	}

	return &artist, nil
}

func (repo *clientConnection) SetArtistCache(ctx context.Context, id int64, artist entity.Artist) error {
	key := fmt.Sprintf(artistDetailKey, id)

	artistString, err := json.Marshal(artist)
	if err != nil {
		return err
	}

	if err := repo.client.Set(ctx, key, artistString, expiration).Err(); err != nil {
		return nil
	}

	return nil
}

func (repo *clientConnection) DeleteArtistCache(ctx context.Context, id int64) error {
	key := fmt.Sprintf(artistDetailKey, id)

	if err := repo.client.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}
