package slash

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/TheDiemer/discord-go-panda/config"
)

func GetSongLink(link string) (success bool, info strings.Builder) {
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
		var plat config.SongLink
		json.Unmarshal([]byte(string(body)), &plat)
		success = true
		info.WriteString(plat.PageUrl)
	}
	return
}
