package slash

import (
	"fmt"

	"github.com/TheDiemer/discord-go-panda/config"

	//	"github.com/go-mysql-org/go-mysql/client"
	//	"github.com/go-mysql-org/go-mysql/mysql"
<<<<<<< HEAD
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

=======
	"strings"
)

>>>>>>> main
type quote struct {
	id     int64
	quote  string
	quoted string
	date   string
	//channel string
}

//func GetQuote(id string, quoted string, conf config.Config) (info strings.Builder, err error) {
func GetQuote(id string, quoted string, conf config.Config) (returned quote, err error) {
	// Lets set our command based on what you got
	var command strings.Builder
	if id != "" {
		command.WriteString("select * from quotes where id = '")
		command.WriteString(id)
		command.WriteString("';")
	} else if quoted != "" {
		command.WriteString("select * from quotes where quoted = '")
		command.WriteString(quoted)
		command.WriteString("';")
	} else {
		command.WriteString("select * from quotes where channel = 0 order by rand() limit 1;")
	}
	// response, tmpErr := dbGet(conf.Database.IP, conf.Database.DB_Username, conf.Database.DB_Password, "karma", command.String())
	response, err := dbGet(conf.Database.IP, conf.Database.DB_Username, conf.Database.DB_Password, "quotes", command.String())

	// We've now got either no quotes, one quote, or a bunch of quotes
	if err != nil {
		fmt.Println(err)
	} else {
		if len(response.Values) > 0 {
			tmpid, _ := response.GetIntByName(0, "id")
			tmpquote, _ := response.GetStringByName(0, "quote")
			tmpquoted, _ := response.GetStringByName(0, "quoted")
			tmpdate, _ := response.GetStringByName(0, "date")
			// tmpchannel, _ := response.GetStringByName(0, "channel")
			returned = quote{
				id:    tmpid,
				quote: tmpquote,

				quoted: tmpquoted,

				date: tmpdate,
				//channel: tmpchannel,
			}
			/*
				info.WriteString("```")
				info.WriteString(quote)
				info.WriteString("```\n -- ")
				info.WriteString(quoted)
				info.WriteString(", ")
				info.WriteString(date)
				info.WriteString(" [")
				info.WriteString(id)
				info.WriteString("]")
			*/
		}
	}
	//if len(response.Values) > 0 || err != nil {
	//	// we got a person back, meaning that alias is in use, or there was an issue with the command itself
	//	if len(response.Values) > 0 {
	//		tmp, _ := response.GetStringByName(0, "person")
	//		err = fmt.Errorf("That alias is already in use!")
	//		info.WriteString("\nThat alias is already in use! It currently points to `")
	//		info.WriteString(tmp)
	//		info.WriteString("`!")
	//	} else {
	//		info.WriteString("\nAn error occurred while reading the current values:\n`")
	//		info.WriteString(err.Error())
	//		info.WriteString("`")
	//	}
	//} else {
	//	fmt.Println("newest here")
	//	// either no error AND no people, so lets store this alias
	//	var writeCommand strings.Builder
	//	writeCommand.WriteString("INSERT INTO alias values('")
	//	writeCommand.WriteString(nick)
	//	writeCommand.WriteString("', '")
	//	writeCommand.WriteString(alias)
	//	writeCommand.WriteString("');")
	//	err = dbWrite(conf.Database.IP, conf.Database.DB_Username, conf.Database.DB_Password, "karma", writeCommand.String())
	//	if err == nil {
	//		info.WriteString("\n`")
	//		info.WriteString(nick)
	//		info.WriteString("` is now known as `")
	//		info.WriteString(alias)
	//		info.WriteString("`!")
	//	} else {
	//		info.WriteString("\nSomething happened while saving the new alias:\n`")
	//		info.WriteString(err.Error())
	//		info.WriteString("`")
	//	}
	//}
	return
}
