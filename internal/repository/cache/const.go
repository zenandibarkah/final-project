package cache

import "time"

const (
	artistKey       = "artist"
	artistDetailKey = "artists:%d"
	albumKey        = "album"
	albumDetailKey  = "album:%d"
	songKey         = "song"
	songDetailKey   = "song:%d"
	expiration      = time.Hour * 1
)
