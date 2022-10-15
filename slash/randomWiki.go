package slash

import (
	"encoding/json"

	"github.com/TheDiemer/discord-go-panda/config"
)

func GetWiki() (wiki config.Wiki, err error) {
	// To get our random wiki, we will just ask wikipedia for a random page!
	var randWikiURL string
	randWikiURL = "https://en.wikipedia.org/api/rest_v1/page/random/summary"

	// Call the starting place and get a body back!
	var body []byte
	body, err = myCall(randWikiURL)
	if err != nil {
		return
	}

	// Unmarshal the response INTO our wiki variable
	json.Unmarshal([]byte(body), &wiki)
	return
}
