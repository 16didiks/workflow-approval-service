package service

import (
	"errors"
	"workflow-approval-service/model"
	"workflow-approval-service/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkflowStepService interface {
	Create(step *model.WorkflowStep) error
	GetByWorkflowID(workflowID string) ([]*model.WorkflowStep, error)
}

type workflowStepService struct {
	db   *gorm.DB
	repo repository.WorkflowStepRepository
}

func NewWorkflowStepService(db *gorm.DB, repo repository.WorkflowStepRepository) WorkflowStepService {
	return &workflowStepService{
		db:   db,
		repo: repo,
	}
}

func (s *workflowStepService) Create(step *model.WorkflowStep) error {
	if step.ID == uuid.Nil {
		step.ID = uuid.New()
	}

	if step.Level <= 0 {
		return errors.New("step level must be greater than zero")
	}

	return s.repo.Create(s.db, step)
}

func (s *workflowStepService) GetByWorkflowID(workflowID string) ([]*model.WorkflowStep, error) {
	wfID, err := uuid.Parse(workflowID)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByWorkflowID(s.db, wfID)
}
