package repository

import (
	"workflow-approval-service/model"

	"gorm.io/gorm"
)

type WorkflowRepository interface {
	Create(db *gorm.DB, workflow *model.Workflow) error
	GetAll(db *gorm.DB) ([]model.Workflow, error)
	GetByID(db *gorm.DB, id string) (*model.Workflow, error)
}

type workflowRepository struct{}

func NewWorkflowRepository() WorkflowRepository {
	return &workflowRepository{}
}

func (r *workflowRepository) Create(db *gorm.DB, workflow *model.Workflow) error {
	return db.Create(workflow).Error
}

func (r *workflowRepository) GetAll(db *gorm.DB) ([]model.Workflow, error) {
	var workflows []model.Workflow
	err := db.Find(&workflows).Error
	return workflows, err
}

func (r *workflowRepository) GetByID(db *gorm.DB, id string) (*model.Workflow, error) {
	var workflow model.Workflow
	err := db.First(&workflow, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &workflow, nil
}
