package handler

type Module struct {
	Name     string
	Commands []*Command
	Access   AccessType // 0 is all, 1 is none, 2 is home, 3 is staff-only, 3 is owner-only
}

type AccessType int

const (
	All AccessType = iota
	None
	Home
	Staff
	Owner
)

func InitModule(name string, access AccessType) *Module {
	return &Module{
		Name:     name,
		Commands: nil,
		Access:   access,
	}
}

func (mdl *Module) AddCommands(cmds ...*Command) *Module {
	for _, command := range cmds {
		command.Module = mdl
	}

	mdl.Commands = append(mdl.Commands, cmds...)
	return mdl
}
