package main

import (
	"github.com/salmonllama/fsbot_go/fsbot"
	"github.com/salmonllama/fsbot_go/lib"
	"os"
)

var (
	config = lib.Config()
)

func main() {
	os.Exit(lifecycle())
}

func lifecycle() int {
	bot := startup()
	if bot != nil {
		bot.Connect()
	}
	return 0
}

func startup() *fsbot.FSBot {
	// Creates and returns a new instance of the bot.
	// Populates listeners and command
	// readies the database
	// Not necessarily in that order
	bot := fsbot.New(config)
	bot.InitModules()

	return bot
}
