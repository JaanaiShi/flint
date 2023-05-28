package model

import (
	"errors"
	"github.com/JaanaiShi/flint/common"
	"gorm.io/gorm"
)

type BaseMapper[T any] struct {
}

func (m BaseMapper[T]) Insert(entity T) (err error) {
	err = common.DB.Create(&entity).Error
	if err != nil {
		return
	}
	return
}

func (m BaseMapper[T]) Delete(entity T) (err error) {
	err = common.DB.Delete(*&entity).Error
	if err != nil {
		return
	}
	return
}

func (m BaseMapper[T]) Update(entity T) (err error) {
	err = common.DB.Updates(*&entity).Error
	if err != nil {
		return
	}
	return
}

func (m BaseMapper[T]) List(entity T) (res []T, err error) {
	err = common.DB.Find(&res, *&entity).Error
	if err != nil {
		return
	}
	return
}

func (m BaseMapper[T]) Detail(entity T) (res T, err error) {
	err = common.DB.First(&res, *&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, err
	}
	return res, nil
}
