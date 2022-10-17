package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"

	"github.com/TheDiemer/discord-go-panda/config"
)

func GetRandomGiph(search string, conf config.Config) string {
	startingUrl := "https://api.giphy.com/v1/gifs/search"
	variables := make(map[string]string)
	variables["api_key"] = conf.Giphy.Key
	variables["q"] = search
	variables["rating"] = "pg-13"
	url := UrlBuilder(startingUrl, variables)

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("experienced error:", err)
		return err.Error()
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return err.Error()
		} else {
			var gif config.Giphy
			json.Unmarshal([]byte(body), &gif)
			chosen := gif.Data[rand.Intn(len(gif.Data))]
			fmt.Println(chosen.Url)
			return chosen.Url
		}
	}
}

func UrlBuilder(startingUrl string, variables map[string]string) string {
	base, err := url.Parse(startingUrl)
	if err != nil {
		return "Failure"
	}

	// Add in our Query Params
	params := url.Values{}
	// Lets loop over our incoming variables
	for key, value := range variables {
		params.Add(key, value)
	}

	// Encoding it onto the end of the url
	base.RawQuery = params.Encode()

	// Return the string of our generated/encoded url
	// fmt.Println(base.String())
	return base.String()

}
