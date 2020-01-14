package handler

type Handler struct {
	Commands []*Command
	Modules  []*Module
}

func (hdlr *Handler) addCommand(cmd *Command) *Command {
	hdlr.Commands = append(hdlr.Commands, cmd)
	return cmd
}

func (hdlr *Handler) GetCommandMap() map[string]*Command {
	cmdMap := make(map[string]*Command)

	for _, cmd := range hdlr.Commands {
		cmdMap[cmd.Name] = cmd
	}

	return cmdMap
}

func (hdlr *Handler) AddModule(mdl *Module) *Module {
	hdlr.Modules = append(hdlr.Modules, mdl)

	for _, cmd := range mdl.Commands {
		hdlr.addCommand(cmd)
	}

	return mdl
}

func (hdlr *Handler) GetModuleMap() map[string]*Module {
	mdlMap := make(map[string]*Module)

	for _, mdl := range hdlr.Modules {
		mdlMap[mdl.Name] = mdl
	}

	return mdlMap
}
