package entities

import (
	"github.com/google/uuid"
	"time"
)

type ID uuid.UUID

type Traceable struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string
	UpdatedBy string
}

type Entity struct {
	ID ID
}

type TraceableEntity struct {
	Entity
	Traceable
}