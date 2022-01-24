package slash

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetSong() (song string, err error) {
	// To get our random song, we will get a bunch of artists from deezer
	// which makes that our...Starting Place!
	var startingPlace string
	startingPlace = "https://api.deezer.com/genre/0/artists"

	// Call the starting place and get a body back!
	var body []byte
	body, err = myCall(startingPlace)
	if err != nil {
		return
	}

	// Define the variable that'll hold all the artists and their data
	var deezer Deezer
	// Unmarshal the response INTO said variable
	json.Unmarshal([]byte(body), &deezer)

	// And now to choose an artist at Random!
	artist := deezer.Data[rand.Intn(len(deezer.Data))]
	fmt.Println(artist)
	tmp := "https://api.deezer.com/artist/" + strconv.Itoa(artist.ID) + "/albums"

	var body2 []byte
	body2, err = myCall(tmp)
	if err != nil {
		return
	}
	// Define the variable that'll hold all the artists and their data
	var artistAlbums Albums
	var listOfAlbums []Album
	// Unmarshal the response INTO said variable
	json.Unmarshal([]byte(body2), &artistAlbums)
	if artistAlbums.Next == "" {
		for _, album := range artistAlbums.Data {
			listOfAlbums = append(listOfAlbums, album)
		}
	} else {
		var albumList []Albums
		albumList = append(albumList, artistAlbums)
		for {
			// Define the variable that'll hold all the artists and their data
			var tmpAlbums Albums
			var body3 []byte
			body3, err = myCall(artistAlbums.Next)
			if err != nil {
				return
			}
			// Unmarshal the response INTO said variable
			json.Unmarshal([]byte(body3), &tmpAlbums)
			albumList = append(albumList, tmpAlbums)
			if tmpAlbums.Next == "" {
				break
			}
			artistAlbums = tmpAlbums
		}

		for _, list := range albumList {
			for _, album := range list.Data {
				listOfAlbums = append(listOfAlbums, album)
			}
		}
	}
	// Now to choose the randomly selected album
	album := listOfAlbums[rand.Intn(len(listOfAlbums))]
	fmt.Println(album)

	var chosenAlbum ChosenAlbum
	var albumData []byte
	albumData, err = myCall(album.Tracklist)
	if err != nil {
		return
	}
	json.Unmarshal([]byte(albumData), &chosenAlbum)

	// And now to choose a song at Random!
	ChosenSong := chosenAlbum.Data[rand.Intn(len(chosenAlbum.Data))]
	fmt.Println(ChosenSong.Title)

	songLinkUrl := "https://api.song.link/v1-alpha.1/links?url="
	songLinkUrl += url.PathEscape(ChosenSong.Link)
	songLinkUrl += "&userCountry=US"

	var songLinkInfo SongLink
	var songLinkData []byte
	songLinkData, err = myCall(songLinkUrl)
	if err != nil {
		return
	}
	json.Unmarshal([]byte(songLinkData), &songLinkInfo)
	song = songLinkInfo.PageUrl
	return
}

func myCall(myURL string) (body []byte, err error) {
	var resp *http.Response
	resp, err = http.Get(myURL)
	if err != nil {
		fmt.Println("issues getting the output back. Currently the url is: ", myURL)
		fmt.Println("and the error is: ", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("issues getting the output back. Currently the url is: ", myURL)
		fmt.Println("and the error is: ", err)
	} else {
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}
	}
	return
}
