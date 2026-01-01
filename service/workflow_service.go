package service

import (
	"errors"

	"workflow-approval-service/model"
	"workflow-approval-service/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrWorkflowNotFound = errors.New("workflow not found")

type WorkflowService interface {
	Create(workflow *model.Workflow) error
	GetAll() ([]model.Workflow, error)
	GetByID(id string) (*model.Workflow, error)
}

type workflowService struct {
	db          *gorm.DB
	workflowRepo repository.WorkflowRepository
}

func NewWorkflowService(db *gorm.DB, repo repository.WorkflowRepository) WorkflowService {
	return &workflowService{
		db:          db,
		workflowRepo: repo,
	}
}

func (s *workflowService) Create(workflow *model.Workflow) error {
	if workflow.Name == "" {
		return errors.New("workflow name is required")
	}
	workflow.ID = uuid.New()
	return s.workflowRepo.Create(s.db, workflow)
}

func (s *workflowService) GetAll() ([]model.Workflow, error) {
	return s.workflowRepo.GetAll(s.db)
}

func (s *workflowService) GetByID(id string) (*model.Workflow, error) {
	return s.workflowRepo.GetByID(s.db, id)
}
