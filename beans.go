package main

import (
	"math/rand"
	"strings"
	"time"
)

func gimmeBeans(user string) (beans strings.Builder) {
	// Define the slice(array) of bean urls
	urls := []string{"https://farm1.static.flickr.com/224/509478598_606e6a436d_o.jpg", "https://i.redd.it/bcykiuff3lz41.jpg", "https://i.redd.it/ne6jt50jwra61.jpg", "https://i.imgur.com/BmrdLMr.jpg"}
	// Go get ONE url, randomly
	bean := randomBean(urls)
	
	// Lets concat the strings now
//	var beans strings.Builder
	beans.WriteString("hey ")
	
	// m.Author.Mention() is the direct person's ID as the proper format for highlighting them
	// based on https://github.com/bwmarrin/discordgo/blob/f36553e31f880f147ae8da1214bf5afca172e326/user.go#L84-L87
	beans.WriteString(user)
	beans.WriteString(", ")
	beans.WriteString(bean)
	return beans
}

// init sets initial values for variables used in the function.
func init() {
        rand.Seed(time.Now().UnixNano())
}

// randomBean returns one of a set of urls. The returned message is selected at random.
func randomBean(urls []string) string {
        return urls[rand.Intn(len(urls))]
}
