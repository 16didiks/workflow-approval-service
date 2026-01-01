package repository

import "gorm.io/gorm"

type TransactionManager interface {
	WithinTransaction(fn func(tx *gorm.DB) error) error
}
