package repository

import (
	"workflow-approval-service/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type workflowStepRepository struct{}

func NewWorkflowStepRepository() WorkflowStepRepository {
	return &workflowStepRepository{}
}

func (r *workflowStepRepository) Create(db *gorm.DB, step *model.WorkflowStep) error {
	return db.Create(step).Error
}

func (r *workflowStepRepository) FindByWorkflowIDAndLevel(db *gorm.DB, workflowID uuid.UUID, level int) (*model.WorkflowStep, error) {
	var step model.WorkflowStep
	err := db.Where("workflow_id = ? AND level = ?", workflowID, level).First(&step).Error
	if err != nil {
		return nil, err
	}
	return &step, nil
}

func (r *workflowStepRepository) FindByWorkflowID(db *gorm.DB, workflowID uuid.UUID) ([]*model.WorkflowStep, error) {
	var steps []*model.WorkflowStep
	err := db.Where("workflow_id = ?", workflowID).Order("level ASC").Find(&steps).Error
	if err != nil {
		return nil, err
	}
	return steps, nil
}
