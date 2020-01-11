package handler

import (
	"github.com/andersfylling/disgord"
	"github.com/salmonllama/fsbot_go/lib"
)

type Command struct {
	Name string
	Description string
	Help string
	Aliases []string
	Permissions []disgord.PermissionBit
	Run func(ctx CommandContext)
}

type CommandContext struct {
	Prefix string
	Args []string
	Command string
	Message *disgord.Message
	Config *lib.Configuration
	Client *disgord.Client
	// Database, when it exists
}

func NewCommand(name string, aliases ...string) *Command {
	return &Command{Name: name, Aliases: aliases}
}

func (cmd *Command) SetDescription(desc string) *Command {
	cmd.Description = desc
	return cmd
}

func (cmd *Command) SetHelp(help string) *Command {
	cmd.Help = help
	return cmd
}

func (cmd *Command) SetPermissions(perms ...disgord.PermissionBit) *Command {
	cmd.Permissions = perms
	return cmd
}

func (cmd *Command) SetOnAction(run func(ctx CommandContext)) *Command {
	cmd.Run = run
	return cmd
}