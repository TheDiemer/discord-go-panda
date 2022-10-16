package config

type RetrievedQuote struct {
	id      int64
	quote   string
	quoted  string
	date    string
	channel string
}

type NewQuote struct {
	quote   string
	quoted  string
	quoter  string
	channel string
}

type Giphy struct {
	Data       []GiphyData `json:"data"`
	Pagination struct {
		TotalCount int `json:"total_count:`
		Count      int `json:"count"`
		Offset     int `json:"offset"`
	} `json:"pagination"`
	Meta struct {
		Status     int    `json:"status"`
		Msg        string `json:"msg"`
		ResponseID string `json:"response_id"`
	} `json:"meta"`
}

type GiphyData struct {
	Type             string `json:"type"`
	ID               string `json:"id"`
	Url              string `json:"url"`
	Slug             string `json:"slug"`
	BitlyGifUrl      string `json:"bitly_gif_url"`
	BitlyUrl         string `json:"bitly_url"`
	EmbedUrl         string `json:"embed_url"`
	Username         string `json:"username"`
	Source           string `json:"source"`
	Title            string `json:"title"`
	Rating           string `json:"rating"`
	ContentUrl       string `json:"content_url"`
	SourceTld        string `json:"source_tld"`
	SourcePostUrl    string `json:"source_post_url"`
	IsSticker        string `json:"is_sticker"`
	ImportDatetime   string `json:"import_datetime"`
	TrendingDatetime string `json:"trending_datetime"`
	Images           struct {
		Original struct {
			Height   string `json:"height"`
			Width    string `json:"width"`
			Size     string `json:"size"`
			Url      string `json:"url"`
			MP4Size  string `json:"mp4_size"`
			MP4      string `json:"mp4"`
			WebpSize string `json:"webp_size"`
			Webp     string `json:"webp"`
			Frames   string `json:"frames"`
			Hash     string `json:"hash"`
		} `json:"original"`
		Downsized struct {
			Height string `json:"height"`
			Width  string `json:"width"`
			Size   string `json:"size"`
			Url    string `json:"url"`
		} `json:"downsized"`
		DownsizedLarge struct {
			Height string `json:"height"`
			Width  string `json:"width"`
			Size   string `json:"size"`
			Url    string `json:"url"`
		} `json:"downsized_large"`
		DownsizedMedium struct {
			Height string `json:"height"`
			Width  string `json:"width"`
			Size   string `json:"size"`
			Url    string `json:"url"`
		} `json:"downsized_medium"`
		DownsizedSmall struct {
			Height string `json:"height"`
			Width  string `json:"width"`
			Size   string `json:"size"`
			Url    string `json:"url"`
		} `json:"downsized_small"`
		DownsizedStill struct {
			Height string `json:"height"`
			Width  string `json:"width"`
			Size   string `json:"size"`
			Url    string `json:"url"`
		} `json:"downsized_still"`
		FixedHeight struct {
			Height   string `json:"height"`
			Width    string `json:"width"`
			Size     string `json:"size"`
			Url      string `json:"url"`
			MP4Size  string `json:"mp4_size"`
			MP4      string `json:"mp4"`
			WebpSize string `json:"webp_size"`
			Webp     string `json:"webp"`
		} `json:"fixed_height"`
		FixedHeightDownsampled struct {
			Height   string `json:"height"`
			Width    string `json:"width"`
			Size     string `json:"size"`
			Url      string `json:"url"`
			WebpSize string `json:"webp_size"`
			Webp     string `json:"webp"`
		} `json:"fixed_height_downsampled"`
		FixedHeightSmall struct {
			Height   string `json:"height"`
			Width    string `json:"width"`
			Size     string `json:"size"`
			Url      string `json:"url"`
			MP4Size  string `json:"mp4_size"`
			MP4      string `json:"mp4"`
			WebpSize string `json:"webp_size"`
			Webp     string `json:"webp"`
		} `json:"fixed_height_small"`
		FixedHeightSmallStill struct {
			Height string `json:"height"`
			Width  string `json:"width"`
			Size   string `json:"size"`
			Url    string `json:"url"`
		} `json:"fixed_height_small_still"`
		FixedHeightStill struct {
			Height string `json:"height"`
			Width  string `json:"width"`
			Size   string `json:"size"`
			Url    string `json:"url"`
		} `json:"fixed_height_still"`
		FixedWidth struct {
			Height   string `json:"height"`
			Width    string `json:"width"`
			Size     string `json:"size"`
			Url      string `json:"url"`
			MP4Size  string `json:"mp4_size"`
			MP4      string `json:"mp4"`
			WebpSize string `json:"webp_size"`
			Webp     string `json:"webp"`
		} `json:"fixed_width"`
		FixedWidthDownsampled struct {
			Height   string `json:"height"`
			Width    string `json:"width"`
			Size     string `json:"size"`
			Url      string `json:"url"`
			WebpSize string `json:"webp_size"`
			Webp     string `json:"webp"`
		} `json:"fixed_width_downsampled"`
		FixedWidthSmall struct {
			Height   string `json:"height"`
			Width    string `json:"width"`
			Size     string `json:"size"`
			Url      string `json:"url"`
			MP4Size  string `json:"mp4_size"`
			MP4      string `json:"mp4"`
			WebpSize string `json:"webp_size"`
			Webp     string `json:"webp"`
		} `json:"fixed_width_small"`
		FixedWidthSmallStill struct {
			Height string `json:"height"`
			Width  string `json:"width"`
			Size   string `json:"size"`
			Url    string `json:"url"`
		} `json:"fixed_width_small_still"`
		FixedWidthStill struct {
			Height string `json:"height"`
			Width  string `json:"width"`
			Size   string `json:"size"`
			Url    string `json:"url"`
		} `json:"fixed_width_still"`
		Looping struct {
			MP4Size string `json:"mp4_size"`
			MP4     string `json:"mp4"`
		} `json:"looping"`
		OriginalStill struct {
			Height string `json:"height"`
			Width  string `json:"width"`
			Size   string `json:"size"`
			Url    string `json:"url"`
		} `json:"original_still"`
		OriginalMP4 struct {
			Height  string `json:"height"`
			Width   string `json:"width"`
			MP4Size string `json:"mp4_size"`
			MP4     string `json:"mp4"`
		} `json:"original_mp4"`
		Preview struct {
			Height string `json:"height"`
			Width  string `json:"width"`
			Size   string `json:"size"`
			Url    string `json:"url"`
		} `json:"preview"`
		PeviewGif struct {
			Height string `json:"height"`
			Width  string `json:"width"`
			Size   string `json:"size"`
			Url    string `json:"url"`
		} `json:"preview_gif"`
		PreviewWebp struct {
			Height string `json:"height"`
			Width  string `json:"width"`
			Size   string `json:"size"`
			Url    string `json:"url"`
		} `json:"preview_webp"`
		FourEightyWStill struct {
			Height string `json:"height"`
			Width  string `json:"width"`
			Size   string `json:"size"`
			Url    string `json:"url"`
		} `json:"480w_still"`
	} `json:"images"`
	User struct {
		AvatarUrl    string `json:"avatar_url"`
		BannerImage  string `json:"banner_image"`
		BannerUrl    string `json:"banner_url"`
		ProfileUrl   string `json:"profile_url"`
		Username     string `json:"username"`
		DisplayName  string `json:"display_name"`
		Description  string `json:"description"`
		InstagramUrl string `json:"instagram_url"`
		WebsiteUrl   string `json:"website_url"`
		IsVerified   bool   `json:"is_verified"`
	} `json:"user"`
	AnalyticsResponsePayload string `json:"analytics_response_payload"`
	Analytics                struct {
		Onload struct {
			Url string `json:"url"`
		}
		Onclick struct {
			Url string `json:"url"`
		}
		Onsent struct {
			Url string `json:"url"`
		}
	} `json:"analytics"`
}

type Wiki struct {
	Type         string `json:"type"`
	Title        string `json:"title"`
	DisplayTitle string `json:"displaytitle"`
	NameSpace    struct {
		ID   int    `json:"id"`
		Text string `json:"text"`
	} `json:"namespace"`
	WikibaseItem string `json:"wikibase_item"`
	Titles       struct {
		Canonical  string `json:"canonical"`
		Normalized string `json:"normalized"`
		Display    string `json:"display"`
	} `json:"titles"`
	PageID    int `json:"pageid"`
	Thumbnail struct {
		Source string `json:"source"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"thumbnail"`
	OriginalImage struct {
		Source string `json:"source"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"originalimage"`
	Lang              string `json:"lang"`
	Dir               string `json:"dir"`
	Revision          string `json:"revision"`
	TID               string `json:"tid"`
	Timestamp         string `json:"timestamp"`
	Description       string `json:"description"`
	DescriptionSource string `json:"description_source"`
	ContentURLs       struct {
		Desktop struct {
			Page      string `json:"page"`
			Revisions string `json:"revisions"`
			Edit      string `json:"edit"`
			Talk      string `json:"talk"`
		} `json:"desktop"`
		Mobile struct {
			Page      string `json:"page"`
			Revisions string `json:"revisions"`
			Edit      string `json:"edit"`
			Talk      string `json:"talk"`
		} `json:"mobile"`
	} `json:"content_urls"`
	Extract     string `json:"extract"`
	ExtractHTML string `json:"extract_html"`
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
	Radio        bool   `json:"radio"`
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
