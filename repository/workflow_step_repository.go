package repository

import (
	"workflow-approval-service/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkflowStepRepository interface {
	Create(db *gorm.DB, step *model.WorkflowStep) error
	FindByWorkflowIDAndLevel(db *gorm.DB, workflowID uuid.UUID, level int) (*model.WorkflowStep, error)
	FindByWorkflowID(db *gorm.DB, workflowID uuid.UUID) ([]*model.WorkflowStep, error)
}
