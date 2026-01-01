package repository

import "gorm.io/gorm"

type TxManager interface {
	WithTransaction(fn func(tx *gorm.DB) error) error
}
