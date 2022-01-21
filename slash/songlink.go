package slash

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Response struct {
	EntityUniqueID     string             `json:"entityUniqueId"`
	UserCountry        string             `json:"userCountry"`
	PageUrl            string             `json:"pageUrl"`
	EntitiesByUniqueID EntitiesByUniqueID `json:"entitiesByUniqueId"`
	LinksByPlatform    LinksByPlatform    `json:"linksByPlatform"`
}

type EntitiesByUniqueID struct {
	ID             string     `json:"id"`
	Type           string     `json:"type"`
	Title          string     `json:"title"`
	ArtistName     string     `json:"artistName"`
	ThumbnailUrl   string     `json:"thumbnailUrl"`
	ThumbnailWidth string     `json:"thumbnailWidth"`
	ApiProvider    string     `json:"apiProvider"`
	Platforms      []struct{} `json:"platforms"`
}

type LinksByPlatform struct {
	Platform Platform
}

type Platform struct {
	Country        string `json:"country"`
	Url            string `json:"url"`
	EntityUniqueID string `json:"entityUniqueId"`
}

func SongLink(link string) (success bool, info strings.Builder) {
	myUrl := "https://api.song.link/v1-alpha.1/links?url="
	myUrl += url.PathEscape(link)
	myUrl += "&userCountry=US"

	resp, err := http.Get(myUrl)
	if err != nil {
		success = false
		info.WriteString("Issues making the initial query: ")
		info.WriteString(err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		success = false
		info.WriteString("Status code returned from querying: ")
		info.WriteString(myUrl)
		info.WriteString("\nThe code returned is: ")
		info.WriteString(string(resp.StatusCode))
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			success = false
			info.WriteString("Error occurred reading the response: ")
			info.WriteString(err.Error())
		}
		var plat Response
		json.Unmarshal([]byte(string(body)), &plat)
		success = true
		info.WriteString(plat.PageUrl)
	}
	return
}
