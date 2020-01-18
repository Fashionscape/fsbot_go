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

// Takes a message and decides if it should be treated as a command or not
func (bot *FSBot) HandleCommand(session disgord.Session, event *disgord.MessageCreate) {
	msg := event.Message
	_, content := bot.separatePrefix(msg.Content)

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

	// Populate the command's context
	ctx := handler.CommandContext{
		Prefix:  bot.Config.DefaultPrefix,
		Args:    nil,
		Command: content,
		Message: msg,
		Client:  bot.Client,
		Config:  &bot.Config,
	}

	// Check if it's an existing command
	for _, cmd := range bot.Handler.Commands {
		for _, alias := range cmd.Aliases {
			if alias == content {
				go cmd.Run(ctx)
			}
		}
	}
}

func (bot *FSBot) HandleImage(session disgord.Session, event *disgord.MessageCreate) {

}

func (bot *FSBot) mdlwImageFilter(event interface{}) interface{} {
	// Filter returns messages that contain an image attachment
	msg := (event.(*disgord.MessageCreate)).Message
	if len(msg.Attachments) > 0 {
		for _, att := range msg.Attachments {
			if lib.IsImage(att.Filename) {
				return event
			}
		}
	}

	return nil
}

func (bot *FSBot) mdlwHasPrefix(event interface{}) interface{} {
	// Filter returns messages that begin with a prefix
	msg := (event.(*disgord.MessageCreate)).Message
	if strings.HasPrefix(msg.Content, bot.Config.DefaultPrefix) {
		return event
	}
	return nil
}

func (bot *FSBot) mdlwIsValidSource(event interface{}) interface{} {
	// Filter blocks bots and blacklisted users and guilds
	msg := (event.(*disgord.MessageCreate)).Message
	//guildId := msg.GuildID.String()
	author := msg.Author

	if author.Bot {
		return nil
	}
	// if user blacklisted
	// if guild blacklisted
	return event
}

func (bot *FSBot) separatePrefix(msg string) (rune, string) {
	r, i := utf8.DecodeRuneInString(msg)
	return r, msg[i:]
}

func (bot *FSBot) hasPermission(member *disgord.Member, command *handler.Command) bool {
	userPerms, err := member.GetPermissions(context.Background(), bot.Client)
	lib.Check(err)
	cmdPerms := command.Permissions

	if len(cmdPerms) == 0 {
		// No command permissions -> Allow usage
		return true
	} else if bot.Config.OwnerID == member.User.ID.String() {
		// If the bot owner executed the command -> Allow usage
		return true
	} else {
		// If the message author has the permission requirement -> Allow usage
		for _, perm := range cmdPerms {
			if perm == userPerms {
				return true
			}
		}
	}

	return false
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

	client.On(disgord.EvtMessageCreate, fsbot.mdlwIsValidSource, fsbot.mdlwImageFilter, fsbot.HandleImage)
	client.On(disgord.EvtMessageCreate, fsbot.mdlwIsValidSource, fsbot.mdlwHasPrefix, fsbot.HandleCommand)

	return fsbot
}
