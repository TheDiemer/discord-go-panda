package slash

import (
	"encoding/json"
)

func GetWiki() (wiki Wiki, err error) {
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

// func myCall(myURL string) (body []byte, err error) {
// 	var resp *http.Response
// 	resp, err = http.Get(myURL)
// 	if err != nil {
// 		fmt.Println("issues getting the output back. Currently the url is: ", myURL)
// 		fmt.Println("and the error is: ", err)
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != 200 {
// 		fmt.Println("issues getting the output back. Currently the url is: ", myURL)
// 		fmt.Println("and the error is: ", err)
// 	} else {
// 		body, err = ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 	}
// 	return
// }
