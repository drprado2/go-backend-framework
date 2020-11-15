package beforeexecute

import "github.com/drprado2/go-backend-framework/pkg/commands"

type Interface interface {
	Execute(command commands.CommandBase, result commands.CommandResult) commands.CommandResult
}

type DefaultBeforeExecuteCommand struct {}

func (*DefaultBeforeExecuteCommand) Execute(command commands.CommandBase, result commands.CommandResult) commands.CommandResult{
	return result
}