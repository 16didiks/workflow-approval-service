package config

import (
	"workflow-approval-service/model"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Workflow{},
		&model.WorkflowStep{},
		&model.Request{},
	)
}
