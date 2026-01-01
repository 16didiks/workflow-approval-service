package repository

import "gorm.io/gorm"

type GormTransactionManager struct {
	db *gorm.DB
}

func NewGormTransactionManager(db *gorm.DB) TransactionManager {
	return &GormTransactionManager{db: db}
}

func (g *GormTransactionManager) WithinTransaction(
	fn func(tx *gorm.DB) error,
) error {
	return g.db.Transaction(fn)
}
