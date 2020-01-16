package fsbot

import (
	"context"
	"github.com/andersfylling/disgord"
	"github.com/salmonllama/fsbot_go/handler"
	"github.com/salmonllama/fsbot_go/lib"
	"github.com/salmonllama/fsbot_go/modules"
	"strings"
	"unicode/utf8"
)

type FSBot struct {
	Client  *disgord.Client
	Config  lib.Configuration
	Handler *handler.Handler
	// TODO: Add Database and Handler to the FSBot struct
}

func (bot *FSBot) isHomeGuild(id string) bool {
	return bot.Config.HomeGuild == id
}

// Takes a message and decided if it should be treated as a command or not
func (bot *FSBot) ProcessCommand(session disgord.Session, evt *disgord.MessageCreate) {
	msg := evt.Message
	// Ignore bot users
	if msg.Author.Bot {
		return
	}

	// Check that it's actually a command
	if len(msg.Content) == 1 {
		return
	}
	if !strings.HasPrefix(msg.Content, bot.Config.DefaultPrefix) {
		return
	}

	ctx := handler.CommandContext{
		Prefix:  bot.Config.DefaultPrefix,
		Args:    nil,
		Command: removePrefix(msg.Content),
		Message: msg,
		Client:  bot.Client,
		Config:  &bot.Config,
	}

	// Check if it's an existing command
	msgContent := removePrefix(msg.Content)
	for _, cmd := range bot.Handler.Commands {
		for _, alias := range cmd.Aliases {
			if alias == msgContent {
				go cmd.Run(ctx)
			}
		}
	}
}

func (bot *FSBot) StoreImage(session disgord.Session, evt *disgord.MessageCreate) {

}

func mdlwImageFilter(evt interface{}) (ret interface{}) {
	// Filter returns messages that contain an image attachment
	msg := (evt.(*disgord.MessageCreate)).Message
	if len(msg.Attachments) > 0 {
		for _, att := range msg.Attachments {
			if lib.IsImage(att.Filename) {
				// return evt
			}
		}
	}

	return nil
}

func removePrefix(msg string) string {
	_, i := utf8.DecodeRuneInString(msg)
	return msg[i:]
}

func (bot *FSBot) InitModules() {
	bot.addModule(modules.ModuleGeneral())
}

func (bot *FSBot) addModule(mdl *handler.Module) *FSBot {
	bot.Handler.AddModule(mdl)
	return bot
}

// Connect opens the connection to discord
func (bot *FSBot) Connect() error {
	err := bot.Client.StayConnectedUntilInterrupted(context.Background())
	lib.Check(err)
	return nil
}

// New creates a new instance of FSBot
func New(config lib.Configuration) *FSBot {
	client := disgord.New(disgord.Config{
		BotToken: config.Token,
	})

	cmd := handler.Handler{}

	fsbot := &FSBot{
		Client:  client,
		Config:  config,
		Handler: &cmd,
	}

	client.On(disgord.EvtMessageCreate, fsbot.ProcessCommand)

	return fsbot
}
