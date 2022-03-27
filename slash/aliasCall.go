package slash

import (
	"fmt"
	"github.com/TheDiemer/discord-go-panda/config"
	"github.com/forPelevin/gomoji"
	"github.com/go-mysql-org/go-mysql/client"
	"github.com/go-mysql-org/go-mysql/mysql"
	"strings"
)

func DeEmoji(sentence string) (replaced strings.Builder) {
	if gomoji.ContainsEmoji(sentence) {
		words := strings.Split(sentence, " ")
		for i, word := range words {
			info, err := gomoji.GetInfo(word)
			if err == nil {
				// This is the condition that means the word was an emoji
				replaced.WriteString(info.Slug)
			} else {
				// And this means it was just a word
				replaced.WriteString(word)
			}
			// Either way we need to reinject spaces
			if i < len(words)-1 {
				replaced.WriteString(" ")
			}
		}
	} else {
		// Since there are no emojis in the sentence, we can return it As is
		replaced.WriteString(sentence)
	}
	return
}

fun GetAlias(name string, conf config.Config) (info strings.Builder, err error) {
}

func MakeAlias(nick string, alias string, conf config.Config) (info strings.Builder, err error) {
	// the Very First thing that should be done is remove any emoji
	tmp := DeEmoji(alias)
	alias = tmp.String()
	var command strings.Builder
	// Next, we should check if that alias points to something else
	command.WriteString("select * from alias where alias = '")
	command.WriteString(alias)
	command.WriteString("';")
	// response, tmpErr := dbGet(conf.Database.IP, conf.Database.DB_Username, conf.Database.DB_Password, "karma", command.String())
	response, err := dbGet(conf.Database.IP, conf.Database.DB_Username, conf.Database.DB_Password, "karma", command.String())
	fmt.Println(response.Values)
	if len(response.Values) > 0 || err != nil {
		// we got a person back, meaning that alias is in use, or there was an issue with the command itself
		if len(response.Values) > 0 {
			tmp, _ := response.GetStringByName(0, "person")
			err = fmt.Errorf("That alias is already in use!")
			info.WriteString("\nThat alias is already in use! It currently points to `")
			info.WriteString(tmp)
			info.WriteString("`!")
		} else {
			info.WriteString("\nAn error occurred while reading the current values:\n`")
			info.WriteString(err.Error())
			info.WriteString("`")
		}
	} else {
		fmt.Println("newest here")
		// either no error AND no people, so lets store this alias
		var writeCommand strings.Builder
		writeCommand.WriteString("INSERT INTO alias values('")
		writeCommand.WriteString(nick)
		writeCommand.WriteString("', '")
		writeCommand.WriteString(alias)
		writeCommand.WriteString("');")
		err = dbWrite(conf.Database.IP, conf.Database.DB_Username, conf.Database.DB_Password, "karma", writeCommand.String())
		if err == nil {
			info.WriteString("\n`")
			info.WriteString(nick)
			info.WriteString("` is now known as `")
			info.WriteString(alias)
			info.WriteString("`!")
		} else {
			info.WriteString("\nSomething happened while saving the new alias:\n`")
			info.WriteString(err.Error())
			info.WriteString("`")
		}
	}
	return
}

// This is to GET data out of the DB, aka read only
func dbGet(ip string, username string, password string, table string, command string) (response *mysql.Result, err error) {
	var conn *client.Conn
	conn, err = client.Connect(ip, username, password, table)
	if err != nil {
		return
	}
	response, err = conn.Execute(command)
	fmt.Println("IM HERE1:")
	fmt.Println(len(response.Values))
	fmt.Println("still HERE")
	if err != nil {
		fmt.Println("IM HERE2:")
		fmt.Println(err)
		fmt.Println("IM HERE3:")
	}
	return
}

// This is to PUT data into the DB, aka write
func dbWrite(ip string, username string, password string, table string, command string) (err error) {
	// Do things with the command
	var conn *client.Conn
	conn, err = client.Connect(ip, username, password, table)
	if err != nil {
		return
	}
	conn.Ping()
	r, tmpErr := conn.Execute(command)
	if tmpErr != nil {
		err = tmpErr
	}
	defer r.Close()
	return
}
