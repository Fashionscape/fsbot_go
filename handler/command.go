package handler

import (
	"github.com/andersfylling/disgord"
)

type Command struct {
	Name        string
	Description string
	Help        string
	Module      *Module
	Aliases     []string
	Permissions disgord.PermissionBit
	Run         func(ctx CommandContext)
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

func (cmd *Command) SetPermission(perms disgord.PermissionBit) *Command {
	cmd.Permissions = perms
	return cmd
}

func (cmd *Command) SetOnAction(run func(ctx CommandContext)) *Command {
	cmd.Run = run
	return cmd
}
