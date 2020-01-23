package lib

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Configuration is the basis for the config
// TODO: Add Twitter?
type Configuration struct {
	Token         string `json:"token"`
	HomeGuild     string `json:"home_guild"`
	LogChannel    string `json:"log_channel"`
	OutfitLogChannel string `json:"outfit_log"`
	ReportChannel string `json:"report_channel"`
	ReportLog     string `json:"report_log"`
	OwnerID       string `json:"owner_id"`
	DefaultPrefix string `json:"default_prefix"`
	Imgur         struct {
		ClientID    string `json:"client_id"`
		BearerToken string `json:"bearer_token"`
	} `json:"imgur"`
	Database struct {
		Hostname string `json:"hostname"`
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
	} `json:"database"`
}

// Config returns the configuration options from the json file
func Config() Configuration {
	file, err := os.Open(LocGet("config.json"))
	Check(err)
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	Check(err)

	var config Configuration
	err = json.Unmarshal(contents, &config)
	Check(err)

	return config
}
