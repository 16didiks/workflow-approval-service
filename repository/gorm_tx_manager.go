package repository

import "gorm.io/gorm"

type GormTxManager struct {
	db *gorm.DB
}

func NewGormTxManager(db *gorm.DB) *GormTxManager {
	return &GormTxManager{db: db}
}

func (g *GormTxManager) WithTransaction(fn func(tx *gorm.DB) error) error {
	return g.db.Transaction(fn)
}
