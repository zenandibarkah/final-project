package entity

type Artist struct {
	ID   int64  `json:"id"`
	Name string `jsong:"name"`
}

type Album struct {
	ID         int64      `json:"id"`
	Title      string     `json:"title"`
	Price      float32    `json:"price"`
	DataArtist DataArtist `json:"Artist"`
}

type Song struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Lyric     string    `json:"lyric"`
	DataAlbum DataAlbum `json:"Album"`
}

type DataArtist struct {
	Artist_id int64  `json:"artist_id"`
	Name      string `json:"name"`
}

type DataAlbum struct {
	Album_id   int64      `json:"album_id"`
	Title      string     `json:"title"`
	DataArtist DataArtist `json:"Artist"`
}
