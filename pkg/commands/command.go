package commands

import (
	"github.com/google/uuid"
	"time"
)

type CommandBase struct {
	CommandId uuid.UUID
	CommandPublishDate time.Time
}

type CommandResult struct {
	CodeResult int
	ResultData interface{}
	ContextData map[string]interface{}
}
