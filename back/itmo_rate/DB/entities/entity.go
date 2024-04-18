package entities

import "gorm.io/gorm"

type Entity interface {
}

func GetById[T Entity](db *gorm.DB, id uint) (ret T, err error) {
	err = db.First(&ret, id).Error
	return
}
