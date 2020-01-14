package modules

import (
	"github.com/salmonllama/fsbot_go/handler"
)

func ModuleGeneral() *handler.Module {
	return handler.InitModule("General", 0).AddCommands(ping(), test())
}

func ping() *handler.Command {
	PingCommand := handler.NewCommand("Ping", "ping")
	PingCommand.SetOnAction(func(ctx handler.CommandContext) {
		go ctx.ReplyText("Pong!")
	})

	return PingCommand
}

func test() *handler.Command {
	TestCommand := handler.NewCommand("Test", "test", "t")
	TestCommand.SetOnAction(func(ctx handler.CommandContext) {
		go ctx.ReplyText("Test confirmed")
	})

	return TestCommand
}
