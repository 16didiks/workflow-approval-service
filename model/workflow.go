package model

import (
	"time"

	"github.com/google/uuid"
)

type Workflow struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"not null;unique"`
	CreatedAt time.Time
}
