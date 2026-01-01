package model

import (
	"time"

	"github.com/google/uuid"
)

type RequestStatus string

const (
	RequestStatusPending  RequestStatus = "PENDING"
	RequestStatusApproved RequestStatus = "APPROVED"
	RequestStatusRejected RequestStatus = "REJECTED"
)

type Request struct {
	ID          uuid.UUID     `gorm:"type:uuid;primaryKey"`
	WorkflowID  uuid.UUID     `gorm:"type:uuid;not null"`
	CurrentStep int           `gorm:"not null"`
	Status      RequestStatus `gorm:"type:varchar(20);not null"`
	Amount      float64       `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
