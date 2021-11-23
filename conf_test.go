package main

import "testing"

const conf_name = "conf.exmpl.yml"
const conf_path = ""
const conf_type = "yaml"

func testConf(t *testing.T) {
	conf, err := readFile(conf_name, conf_path, conf_type)
	if err != nil {
		t.Errorf("Error loading the config: '%q'", err)
	}

	expected := Config{}
	expected.Discord.Token = "DISCORD_BOT_TOKEN"
	expected.Discord.Send_time = "16"
	expected.Discord.Owner = "DISCORD_USER_ID"
	expected.Database.IP = "localhost"
	expected.Database.DB_Username = "DB_MASTER_USERNAME"
	expected.Database.DB_Password = "DB_MASTER_PASSWORD"
	expected.Database.DL_Username = "ALT_DB_USERNAME"
	expected.Database.DL_Password = "ALT_DB_PASSWORD"

	if conf != expected {
		t.Errorf("expected '%q' but got '%q'", expected, conf)
	}
}
