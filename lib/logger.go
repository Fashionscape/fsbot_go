package lib

import (
	"fmt"
	"github.com/andersfylling/disgord"
	"github.com/salmonllama/fsbot_go/fsbot"
	"github.com/salmonllama/fsbot_go/handler"
	"time"
)

type Logger struct {
	DiscordChannel   string
	DefaultToDiscord bool
	Bot              *fsbot.FSBot
}

type LogType int // 0 is Info, 1 is Warn, 2 is Err

const (
	Info LogType = iota
	Warn
	Err
)

var (
	base = "[" + time.Now().String() + "]"

	infoString = base + " INFO: "
	infoEmbed  = handler.EmbedBuilder{
		Title: "LOGGER: INFO",
		Color: 16777215,
	}

	warnString = base + " WARN: "
	warnEmbed  = handler.EmbedBuilder{
		Title: "LOGGER: WARN",
		Color: 0,
	}

	errString = base + " ERR: "
	errEmbed  = handler.EmbedBuilder{
		Title: "LOGGER: ERR",
		Color: 0,
	}
)

func (log *Logger) logToConsole(content string) {
	fmt.Println(content)
}

func (log *Logger) logToDiscord(content *disgord.Embed) {

}

func (log *Logger) handleLogging(content string, logType LogType) {
	switch logType {
	case Info:
		if log.DefaultToDiscord {
			embed := infoEmbed.SetDescription(content)
			log.logToDiscord(embed.Build())
		} else {
			s := infoString + content
			log.logToConsole(s)
		}
	case Warn:
		if log.DefaultToDiscord {
			embed := warnEmbed.SetDescription(content)
			log.logToDiscord(embed.Build())
		} else {
			s := warnString + content
			log.logToConsole(s)
		}
	case Err:
		if log.DefaultToDiscord {
			embed := errEmbed.SetDescription(content)
			log.logToDiscord(embed.Build())
		} else {
			s := errString + content
			log.logToConsole(s)
		}
	}
}

func (log *Logger) Info(content string) {

}

func (log *Logger) Warn(content string) {

}

func (log *Logger) Err(content string) {

}
