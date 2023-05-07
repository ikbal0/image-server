package transaction

import (
	"gorm.io/gorm"
)

type TransactionImpl struct {
	db *gorm.DB
}

func NewTransaction(db *gorm.DB) *TransactionImpl {
	tx := &TransactionImpl{db: db}
	return tx
}
