package slash

type Deezer struct {
	Data []Artist `json:"data"`
}

type Artist struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Picture      string `json:"picture"`
	PictureSmall string `json:"picture_small"`
	PictureMed   string `json:"picture_medium"`
	PictureBig   string `json:"picture_big"`
	PictureXL    string `json:"picture_xl"`
	radio        bool   `json:"radio"`
	TrackList    string `json:"tracklist"`
	Type         string `json:"type"`
}

type Albums struct {
	Data  []Album `json:"data"`
	Total int     `json:"total"`
	Prev  string  `json:"prev"`
	Next  string  `json:"next"`
}

type Album struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Cover       string `json:"cover"`
	CoverSmall  string `json:"cover_small"`
	CoverMed    string `json:"cover_medium"`
	CoverBig    string `json:"cover_big"`
	CoverXL     string `json:"cover_xl"`
	MD5Image    string `json:"md5_image"`
	GenreID     int    `json:"genre_id"`
	Fans        int    `json:"fans"`
	ReleaseDate string `json:"release_date"`
	RecordType  string `json:"record_type"`
	Tracklist   string `json:"tracklist"`
	Explicit    bool   `json:"explicit_lyrics"`
	Type        string `json:"album"`
}

type ChosenAlbum struct {
	Data  []Song `json:"data"`
	Total int    `json:"total"`
}

type Song struct {
	ID                    int    `json:"id"`
	Readable              bool   `json:"readable"`
	Title                 string `json:"title"`
	TitleShort            string `json:"title_short"`
	TitleVersion          string `json:"title_version"`
	Isrc                  string `json:"isrc"`
	Link                  string `json:"link"`
	Duration              int    `json:"duration"`
	TrackPosition         int    `json:"track_position"`
	DiskNumber            int    `json:"disk_number"`
	Rank                  int    `json:"rank"`
	ExplicitLyrics        bool   `json:"explicit_lyrics"`
	ExplicitContentLyrics int    `json:"explicit_Content_lyrics"`
	ExplicitContentCover  int    `json:"explicit_content_cover"`
	Preview               string `json:"preview"`
	MD5Image              string `json:"md5_image"`
	Artist                Artist `json:"artist"`
}

type SongLink struct {
	EntityUniqueID     string             `json:"entityUniqueId"`
	UserCountry        string             `json:"userCountry"`
	PageUrl            string             `json:"pageUrl"`
	EntitiesByUniqueID EntitiesByUniqueID `json:"entitiesByUniqueId"`
	LinksByPlatform    LinksByPlatform    `json:"linksByPlatform"`
}

type EntitiesByUniqueID struct {
	ID             string `json:"id"`
	Type           string `json:"type"`
	Title          string `json:"title"`
	ArtistName     string `json:"artistName"`
	ThumbnailUrl   string `json:"thumbnailUrl"`
	ThumbnailWidth string `json:"thumbnailWidth"`
	ApiProvider    string `json:"apiProvider"`
	Platforms      []struct{}
}

type LinksByPlatform struct {
	Platform *Platform
}

type Platform struct {
	Country        string `json:"country"`
	Url            string `json:"url"`
	EntityUniqueID string `json:"entityUniqueId"`
}
