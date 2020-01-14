package handler

import (
	"github.com/andersfylling/disgord"
	"time"
)

type EmbedBuilder struct {
	Title       string
	Type        string
	Description string
	URL         string
	Timestamp   time.Time
	Color       int
	Footer      *disgord.EmbedFooter
	Image       *disgord.EmbedImage
	Thumbnail   *disgord.EmbedThumbnail
	Video       *disgord.EmbedVideo
	Provider    *disgord.EmbedProvider
	Author      *disgord.EmbedAuthor
	Fields      []*disgord.EmbedField
}

func NewRichEmbed(title string) *EmbedBuilder {
	// Probably won't even end up using this in favor of an expanded declaration
	return &EmbedBuilder{Title: title}
}

func (bldr *EmbedBuilder) SetFooter(text string) *EmbedBuilder {
	footer := disgord.EmbedFooter{
		Lockable: disgord.Lockable{},
		Text:     text,
	}
	bldr.Footer = &footer
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

func (bldr *EmbedBuilder) SetProvider(name string, url string) *EmbedBuilder {
	provider := disgord.EmbedProvider{
		Lockable: disgord.Lockable{},
		Name:     name,
		URL:      url,
	}
	bldr.Provider = &provider
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
		Footer:      bldr.Footer,
		Image:       bldr.Image,
		Thumbnail:   bldr.Thumbnail,
		Video:       bldr.Video,
		Provider:    bldr.Provider,
		Author:      bldr.Author,
		Fields:      bldr.Fields,
	}
}
