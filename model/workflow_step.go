package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type WorkflowStep struct {
	ID         uuid.UUID      `gorm:"type:uuid;primaryKey"`
	WorkflowID uuid.UUID      `gorm:"type:uuid;not null;index;uniqueIndex:idx_workflow_level"`
	Level      int            `gorm:"not null;uniqueIndex:idx_workflow_level"`
	Actor      string         `gorm:"not null"`
	Conditions datatypes.JSON `gorm:"type:jsonb"`
}
