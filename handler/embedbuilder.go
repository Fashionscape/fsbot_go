package handler

import (
	"github.com/andersfylling/disgord"
)

type EmbedBuilder struct {
	Title       string
	Description string
	URL         string
	Color       int
	Image       *disgord.EmbedImage
	Thumbnail   *disgord.EmbedThumbnail
	Video       *disgord.EmbedVideo
	Author      *disgord.EmbedAuthor
	Fields      []*disgord.EmbedField
}

func NewRichEmbed(title string) *EmbedBuilder {
	// Probably won't even end up using this in favor of an expanded declaration
	return &EmbedBuilder{Title: title}
}

func (bldr *EmbedBuilder) SetDescription(desc string) *EmbedBuilder {
	bldr.Description = desc
	return bldr
}

func (bldr *EmbedBuilder) SetImage(url string) *EmbedBuilder {
	image := disgord.EmbedImage{
		Lockable: disgord.Lockable{},
		URL:      url,
	}
	bldr.Image = &image
	return bldr
}

func (bldr *EmbedBuilder) SetThumbnail(url string) *EmbedBuilder {
	thumbnail := disgord.EmbedThumbnail{
		Lockable: disgord.Lockable{},
		URL:      url,
	}
	bldr.Thumbnail = &thumbnail
	return bldr
}

func (bldr *EmbedBuilder) SetVideo(url string) *EmbedBuilder {
	video := disgord.EmbedVideo{
		Lockable: disgord.Lockable{},
		URL:      url,
	}
	bldr.Video = &video
	return bldr
}

func (bldr *EmbedBuilder) SetAuthor(user *disgord.User) *EmbedBuilder {
	author := disgord.EmbedAuthor{
		Lockable: disgord.Lockable{},
		Name:     user.Username,
		IconURL:  user.Avatar,
	}
	bldr.Author = &author
	return bldr
}

func (bldr *EmbedBuilder) AddField(name string, value string, inline bool) *EmbedBuilder {
	field := disgord.EmbedField{
		Lockable: disgord.Lockable{},
		Name:     name,
		Value:    value,
		Inline:   inline,
	}
	bldr.Fields = append(bldr.Fields, &field)
	return bldr
}

func (bldr *EmbedBuilder) AddFields(fields ...*disgord.EmbedField) *EmbedBuilder {
	bldr.Fields = fields
	return bldr
}

func (bldr *EmbedBuilder) Build() *disgord.Embed {
	return &disgord.Embed{
		Lockable:    disgord.Lockable{},
		Title:       bldr.Title,
		Type:        "rich",
		Description: bldr.Description,
		URL:         bldr.URL,
		Timestamp:   disgord.Time{},
		Color:       bldr.Color,
		Footer: &disgord.EmbedFooter{
			Lockable: disgord.Lockable{},
			Text:     "FSBot {version}",
		},
		Image:     bldr.Image,
		Thumbnail: bldr.Thumbnail,
		Video:     bldr.Video,
		Provider: &disgord.EmbedProvider{
			Lockable: disgord.Lockable{},
			Name:     "",
			URL:      "",
		},
		Author: bldr.Author,
		Fields: bldr.Fields,
	}
}
