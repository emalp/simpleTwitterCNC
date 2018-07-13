package main

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Config is the configuration taken from config.json file.
type Config struct {
	ConsumerKey    string `json:"ConsumerKey"`
	ConsumerSecret string `json:"ConsumerSecret"`
	AccessToken    string `json:"AccessToken"`
	ATSecret       string `json:"ATSecret"`
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func importCredentials() Config {
	var configuration Config

	data, readErr := ioutil.ReadFile("./config.json")
	checkError(readErr)

	jsonErr := json.Unmarshal(data, &configuration)
	checkError(jsonErr)

	return configuration
}

func initializeController() string {
	conf := importCredentials()

	config := oauth1.NewConfig(conf.ConsumerKey, conf.ConsumerSecret)
	token := oauth1.NewToken(conf.AccessToken, conf.ATSecret)
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	currentUserAccountParameters := &twitter.AccountVerifyParams{
		IncludeEmail:    twitter.Bool(true),
		IncludeEntities: twitter.Bool(true),
		SkipStatus:      twitter.Bool(false)}

	profileData, _, err := client.Accounts.VerifyCredentials(currentUserAccountParameters) // returns profileData, response, error

	checkError(err)
	currentUserTimelineParameters := twitter.UserTimelineParams{
		UserID: profileData.ID}

	tweets, _, error := client.Timelines.UserTimeline(&currentUserTimelineParameters) // _ = resp

	checkError(error)

	// tweets[0] contains the latest tweet
	if len(tweets) > 0 {
		latestTweet := tweets[0]
		return latestTweet.Text
	} else {
		return ""
	}
}

func initialize() {
	var lastCommand string
	for range time.Tick(time.Second * 5) {
		command := initializeController()
		if command != "" {
			if lastCommand == command {
				// do nothing
			} else {
				splitAndRunCommand(command)
			}
		}

		lastCommand = command
	}
}

func main() {
	//fmt.Println("\nStarting timer")

	initialize()
	//splitAndRunCommand("emalp")
}
