package service

import (
	"encoding/json"
	"workflow-approval-service/model"
	"workflow-approval-service/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RequestService interface {
	Create(req *model.Request) error
	GetByID(id string) (*model.Request, error)
	Approve(id string) error
	Reject(id string) error
}

type requestService struct {
	db               *gorm.DB
	requestRepo      repository.RequestRepository
	workflowStepRepo repository.WorkflowStepRepository
}

func NewRequestService(
	db *gorm.DB,
	requestRepo repository.RequestRepository,
	workflowStepRepo repository.WorkflowStepRepository,
) RequestService {
	return &requestService{
		db:               db,
		requestRepo:      requestRepo,
		workflowStepRepo: workflowStepRepo,
	}
}

func (s *requestService) Create(req *model.Request) error {
	if req.Amount <= 0 {
		return ErrRequestAlreadyProcessed
	}

	req.ID = uuid.New()
	req.Status = model.RequestStatusPending
	req.CurrentStep = 1

	return s.requestRepo.Create(s.db, req)
}

func (s *requestService) GetByID(id string) (*model.Request, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, ErrRequestNotFound
	}

	req, err := s.requestRepo.FindByID(s.db, uuidID)
	if err != nil {
		return nil, ErrRequestNotFound
	}

	return req, nil
}

func (s *requestService) Approve(id string) error {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return ErrRequestNotFound
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		req, err := s.requestRepo.FindByIDForUpdate(tx, uuidID)
		if err != nil {
			return ErrRequestNotFound
		}

		if req.Status != model.RequestStatusPending {
			return ErrRequestAlreadyProcessed
		}

		step, err := s.workflowStepRepo.FindByWorkflowIDAndLevel(tx, req.WorkflowID, req.CurrentStep)
		if err != nil {
			return err
		}

		var minAmount float64
		if step != nil && step.Conditions != nil {
			var cond map[string]interface{}
			if err := json.Unmarshal(step.Conditions, &cond); err == nil {
				if v, ok := cond["min_amount"]; ok {
					switch val := v.(type) {
					case float64:
						minAmount = val
					case int:
						minAmount = float64(val)
					case int64:
						minAmount = float64(val)
					}
				}
			}
		}

		if req.Amount >= minAmount {
			req.CurrentStep++
			nextStep, _ := s.workflowStepRepo.FindByWorkflowIDAndLevel(tx, req.WorkflowID, req.CurrentStep)
			if nextStep == nil {
				req.Status = model.RequestStatusApproved
			}
		} else {
			req.Status = model.RequestStatusApproved
		}

		return s.requestRepo.Update(tx, req)
	})
}

func (s *requestService) Reject(id string) error {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return ErrRequestNotFound
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		req, err := s.requestRepo.FindByIDForUpdate(tx, uuidID)
		if err != nil {
			return ErrRequestNotFound
		}

		if req.Status != model.RequestStatusPending {
			return ErrRequestAlreadyProcessed
		}

		req.Status = model.RequestStatusRejected
		return s.requestRepo.Update(tx, req)
	})
}
