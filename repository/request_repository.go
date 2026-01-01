package repository

import (
	"workflow-approval-service/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// RequestRepository interface
type RequestRepository interface {
	Create(db *gorm.DB, req *model.Request) error
	FindByID(db *gorm.DB, id uuid.UUID) (*model.Request, error)
	FindByIDForUpdate(db *gorm.DB, id uuid.UUID) (*model.Request, error)
	Update(db *gorm.DB, req *model.Request) error
}

// requestRepository implements RequestRepository
type requestRepository struct{}

// NewRequestRepository constructor
func NewRequestRepository() RequestRepository {
	return &requestRepository{}
}

// Create a new request
func (r *requestRepository) Create(db *gorm.DB, req *model.Request) error {
	return db.Create(req).Error
}

// FindByID fetch request by ID
func (r *requestRepository) FindByID(db *gorm.DB, id uuid.UUID) (*model.Request, error) {
	var req model.Request
	err := db.Where("id = ?", id).First(&req).Error
	if err != nil {
		return nil, err
	}
	return &req, nil
}

// FindByIDForUpdate fetch request by ID with row lock
func (r *requestRepository) FindByIDForUpdate(db *gorm.DB, id uuid.UUID) (*model.Request, error) {
	var req model.Request
	err := db.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", id).First(&req).Error
	if err != nil {
		return nil, err
	}
	return &req, nil
}

// Update request
func (r *requestRepository) Update(db *gorm.DB, req *model.Request) error {
	return db.Save(req).Error
}
