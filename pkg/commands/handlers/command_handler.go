package handlers

import (
	"github.com/drprado2/go-backend-framework/pkg/commands"
	"github.com/drprado2/go-backend-framework/pkg/commands/beforeexecute"
)

type CommandHandlerInterface interface {
	Handle(command commands.CommandBase, result commands.CommandResult) commands.CommandResult
}

type CommandHandlerExecutorInterface interface {
	ExecuteCommand(command *commands.CommandBase, handler *CommandHandlerInterface) commands.CommandResult
}

type CommandHandlerExecutor struct {
	beforeExecuteActions []*beforeexecute.Interface
}