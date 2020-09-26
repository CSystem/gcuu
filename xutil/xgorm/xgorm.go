package xgorm

import (
	"github.com/jinzhu/gorm"
)

// http://hopehook.com/2017/08/21/golang_transaction/
func Transact(db *gorm.DB, txFunc func(*gorm.DB) error) (err error) {
	tx := db.Begin()
	err = tx.Error

	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	return txFunc(tx)
}
