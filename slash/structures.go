package slash

type Wiki struct {
	Type              string          `json:"type"`
	Title             string          `json:"title"`
	DisplayTitle      string          `json:"displaytitle"`
	NameSpace         []Namespace     `json:"namespace"`
	WikibaseItem      string          `json:"wikibase_item"`
	Titles            []TitlesList    `json:"titles"`
	PageID            int             `json:"pageid"`
	ThumbnailInfo     []Thumbnail     `json:"thumbnail"`
	OriginalImageInfo []OriginalImage `json:"originalimage"`
	Lang              string          `json:"lang"`
	Dir               string          `json:"dir"`
	Revision          string          `json:"revision"`
	TID               string          `json:"tid"`
	Timestamp         string          `json:"timestamp"`
	Description       string          `json:"description"`
	DescriptionSource string          `json:"description_source"`
	ContentURLs       []WikiURLs      `json:"content_urls"`
	Extract           string          `json:"extract"`
	ExtractHTML       string          `json:"extract_html"`
}

type Namespace struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type TitlesList struct {
	Canonical  string `json:"canonical"`
	Normalized string `json:"normalized"`
	Display    string `json:"display"`
}

type Thumbnail struct {
	Source string `json:"source"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type OriginalImage struct {
	Source string `json:"source"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type WikiURLs struct {
	Desktop []URLs `json:"desktop"`
	Mobile  []URLs `json:"mobile"`
}

type URLs struct {
	Page      string `json:"page"`
	Revisions string `json:"revisions"`
	Edit      string `json:"edit"`
	Talk      string `json:"talk"`
}

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
