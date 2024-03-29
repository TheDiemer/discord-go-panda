package slash

import (
	"fmt"

	"github.com/TheDiemer/discord-go-panda/config"

	"strconv"
	"strings"
)

func AddQuote(newQuote config.NewQuote) (info strings.Builder, err error) {
	var command strings.Builder
	command.WriteString("INSERT INTO quotes (quoted, quote, quoter, channel) values ('")
	command.WriteString(newQuote.Quoted)
	command.WriteString("', '")
	command.WriteString(newQuote.Quote)
	command.WriteString("', '")
	command.WriteString(newQuote.Quoter)
	command.WriteString("', '")
	command.WriteString(newQuote.Channel)
	command.WriteString("');")
	fmt.Println(command.String())
	err = dbWrite(conf.Database.IP, conf.Database.DB_Username, conf.Database.DB_Password, "quotes", command.String())
	if err == nil {
		var getCommand strings.Builder
		getCommand.WriteString("select id from quotes where quote = '")
		getCommand.WriteString(newQuote.Quote)
		getCommand.WriteString("';")
		response, err2 := dbGet(conf.Database.IP, conf.Database.DB_Username, conf.Database.DB_Password, "quotes", getCommand.String())
		if err2 == nil {
			info.WriteString("Quote successfully added!\nNew quote id: **")
			tmpid, _ := response.GetIntByName(0, "id")
			info.WriteString(strconv.FormatInt(tmpid, 10))
			info.WriteString("**.")
		} else {
			info.WriteString("Quote was successfully added, but I failed to identify its number...\n**Sorry!**\n¯\\_(ツ)_/¯")
		}
	} else {
		info.WriteString("Quote was not added. Please see the explanation:\n`")
		info.WriteString(err.Error())
		info.WriteString("\n`.")
	}
	return
}

func GetQuote(id string, quoted string, conf config.Config) (returned config.RetrievedQuote, err error) {
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
			tmpchannel, _ := response.GetStringByName(0, "channel")
			returned.ID, returned.Quote, returned.Quoted, returned.Date, returned.Channel = tmpid, tmpquote, tmpquoted, tmpdate, tmpchannel
		}
	}
	return
}
