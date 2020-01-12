package handler

import (
	"context"
	"github.com/andersfylling/disgord"
	"github.com/salmonllama/fsbot_go/lib"
)

type CommandContext struct {
	Prefix string
	Args []string
	Command string
	Message *disgord.Message
	Config *lib.Configuration
	Client *disgord.Client
	// Database, when it exists
}

func (ctx *CommandContext) ReplyText(text string) (msg *disgord.Message) {
	msg, err := ctx.Message.Reply(context.Background(), ctx.Client, text)
	lib.Check(err)
	return
}

func (ctx *CommandContext) ReplyEmbed(embed *disgord.Embed) (msg *disgord.Message) {
	msg, err := ctx.Message.Reply(context.Background(), ctx.Client, embed)
	lib.Check(err)
	return
}

func (ctx *CommandContext) GetAuthorID() (id string) {
	id = ctx.Message.Author.ID.String()
	return
}

func (ctx *CommandContext) GetAuthorMention() (mention string) {
	mention = ctx.Message.Author.Mention()
	return
}

func (ctx *CommandContext) GetChannel() (channel *disgord.Channel) {
	channel, err := ctx.Client.GetChannel(context.Background(), ctx.Message.ChannelID)
	lib.Check(err)
	return
}

func (ctx *CommandContext) GetGuild() (guild *disgord.Guild) {
	guild, err := ctx.Client.GetGuild(context.Background(), ctx.Message.GuildID)
	lib.Check(err)
	return
}

func (ctx *CommandContext) GetMentionedUsers() (users []*disgord.User) {
	users = ctx.Message.Mentions
	return
}

func (ctx *CommandContext) GetMentionedRoles() (roles []*disgord.Role) {
	snowflakes := ctx.Message.MentionRoles
	roles = make([]*disgord.Role, len(snowflakes))
	guild := ctx.GetGuild()
	for _, sf := range snowflakes {
		role, err := guild.Role(sf)
		lib.Check(err)
		roles = append(roles, role)
	}
	return
}

func (ctx *CommandContext) GetMentionedChannels() (channels []*disgord.Channel) {
	mChannels := ctx.Message.MentionChannels
	guild := ctx.GetGuild()
	channels = make([]*disgord.Channel, len(mChannels))
	for _, chl := range mChannels {
		c, err := guild.Channel(chl.ID)
		lib.Check(err)
		channels = append(channels, c)
	}
	return
}
