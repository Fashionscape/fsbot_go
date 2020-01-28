package lib

import (
	"context"
	"fmt"
	"github.com/andersfylling/disgord"
	"time"
)

type Logger struct {
	DiscordChannel   string
	DefaultToDiscord bool
	Client           *disgord.Client
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
	infoEmbed  = EmbedBuilder{
		Title: "LOGGER: INFO",
		Color: 0xFF8000,
	}

	warnString = base + " WARN: "
	warnEmbed  = EmbedBuilder{
		Title: "LOGGER: WARN",
		Color: 0xFFFF00,
	}

	errString = base + " ERR: "
	errEmbed  = EmbedBuilder{
		Title: "LOGGER: ERR",
		Color: 0xFF0100,
	}
)

func (log *Logger) logToConsole(content string) {
	fmt.Println(content)
}

func (log *Logger) logToDiscord(content *disgord.Embed) {
	sf := disgord.ParseSnowflakeString(log.DiscordChannel)
	channel, err := log.Client.GetChannel(context.Background(), sf)
	if err != nil {
		log.logToConsole(err.Error())
		panic(err)
	}
	msg := &disgord.Message{
		Embeds: []*disgord.Embed{content},
	}

	_, err = channel.SendMsg(context.Background(), log.Client, msg)
	if err != nil {
		log.logToConsole(err.Error())
	}
}

func (log *Logger) handleLogging(content string, logType LogType) {
	switch logType {
	case Info:
		if log.DefaultToDiscord {
			embed := infoEmbed.SetDescription(content)
			go log.logToDiscord(embed.Build())
		} else {
			s := infoString + content
			log.logToConsole(s)
		}
	case Warn:
		if log.DefaultToDiscord {
			embed := warnEmbed.SetDescription(content)
			go log.logToDiscord(embed.Build())
		} else {
			s := warnString + content
			log.logToConsole(s)
		}
	case Err:
		if log.DefaultToDiscord {
			embed := errEmbed.SetDescription(content)
			go log.logToDiscord(embed.Build())
		} else {
			s := errString + content
			log.logToConsole(s)
		}
	}
}

func (log *Logger) Info(content string) {
	log.handleLogging(content, Info)
}

func (log *Logger) Warn(content string) {
	log.handleLogging(content, Warn)
}

func (log *Logger) Err(content string) {
	log.handleLogging(content, Err)
}
