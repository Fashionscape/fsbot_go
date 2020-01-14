package handler

type Module struct {
	Name     string
	Commands []*Command
	Access   AccessType // 0 is all, 1 is none, 2 is home
}

type AccessType int

const (
	All AccessType = iota
	None
	Home
)

func InitModule(name string, access AccessType) *Module {
	return &Module{
		Name:     name,
		Commands: nil,
		Access:   access,
	}
}

func (mdl *Module) AddCommands(cmd ...*Command) *Module {
	mdl.Commands = append(mdl.Commands, cmd...)
	return mdl
}
